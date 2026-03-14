package user

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"clarity-gym/internal/auth"
)

type ProfileUpdate struct {
	Name            string `json:"name"`
	Bio             string `json:"bio"`
	AvatarURL       string `json:"avatar_url"`
	Specialty       string `json:"specialty"`
	ExperienceYears int    `json:"experience_years"`
}

func (h *Handler) GetProfile(w http.ResponseWriter, r *http.Request) {
	userID, _ := r.Context().Value(auth.UserIDKey).(string)

	type Profile struct {
		ID              string `json:"id"`
		Name            string `json:"name"`
		Email           string `json:"email"`
		Role            string `json:"role"`
		Bio             string `json:"bio"`
		AvatarURL       string `json:"avatar_url"`
		Specialty       string `json:"specialty"`
		ExperienceYears int    `json:"experience_years"`
	}

	var p Profile
	err := h.DB.QueryRow(`
		SELECT id, name, email, role,
		COALESCE(bio, ''), COALESCE(avatar_url, ''),
		COALESCE(specialty, ''), COALESCE(experience_years, 0)
		FROM users WHERE id = $1
	`, userID).Scan(&p.ID, &p.Name, &p.Email, &p.Role, &p.Bio, &p.AvatarURL, &p.Specialty, &p.ExperienceYears)

	if err != nil {
		http.Error(w, "User negasit", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func (h *Handler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	userID, _ := r.Context().Value(auth.UserIDKey).(string)

	var update ProfileUpdate
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		http.Error(w, "Date invalide", http.StatusBadRequest)
		return
	}

	_, err := h.DB.Exec(`
		UPDATE users SET
			name = $1,
			bio = $2,
			avatar_url = $3,
			specialty = $4,
			experience_years = $5
		WHERE id = $6
	`, update.Name, update.Bio, update.AvatarURL, update.Specialty, update.ExperienceYears, userID)

	if err != nil {
		http.Error(w, "Eroare server", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func (h *Handler) UploadAvatar(w http.ResponseWriter, r *http.Request) {
	userID, _ := r.Context().Value(auth.UserIDKey).(string)

	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("avatar")
	if err != nil {
		http.Error(w, "Fisier invalid", http.StatusBadRequest)
		return
	}
	defer file.Close()

	ext := filepath.Ext(handler.Filename)
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".webp" {
		http.Error(w, "Format invalid. Acceptam: jpg, png, webp", http.StatusBadRequest)
		return
	}

	// Sterge poza veche
	oldFiles, _ := filepath.Glob(filepath.Join("uploads/avatars", userID+".*"))
	for _, f := range oldFiles {
		os.Remove(f)
	}

	filename := userID + ext
	savePath := filepath.Join("uploads/avatars", filename)

	dst, err := os.Create(savePath)
	if err != nil {
		http.Error(w, "Eroare la salvare", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	io.Copy(dst, file)

	avatarURL := "/uploads/avatars/" + filename
	_, err = h.DB.Exec("UPDATE users SET avatar_url = $1 WHERE id = $2", avatarURL, userID)
	if err != nil {
		http.Error(w, "Eroare DB", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"avatar_url": avatarURL})
}
