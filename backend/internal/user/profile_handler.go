package user

import (
	"encoding/json"
	"net/http"

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
