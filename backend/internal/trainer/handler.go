package trainer

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Handler struct {
	DB *sql.DB
}

type Trainer struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	Bio             string `json:"bio"`
	Specialty       string `json:"specialty"`
	ExperienceYears int    `json:"experience_years"`
	AvatarURL       string `json:"avatar_url"`
}

func (h *Handler) GetAllTrainers(w http.ResponseWriter, r *http.Request) {
	rows, err := h.DB.Query(
		`SELECT id, name, email, COALESCE(bio, ''), COALESCE(specialty, ''), 
		COALESCE(experience_years, 0), COALESCE(avatar_url, '') 
		FROM users WHERE role = 'trainer' AND is_suspended = false 
		ORDER BY created_at DESC`,
	)
	if err != nil {
		http.Error(w, "Eroare server", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	trainers := []Trainer{}
	for rows.Next() {
		var t Trainer
		if err := rows.Scan(&t.ID, &t.Name, &t.Email, &t.Bio, &t.Specialty, &t.ExperienceYears, &t.AvatarURL); err != nil {
			continue
		}
		trainers = append(trainers, t)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trainers)
}
