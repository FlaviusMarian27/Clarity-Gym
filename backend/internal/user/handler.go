package user

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	DB *sql.DB
}

type User struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Role        string `json:"role"`
	IsSuspended bool   `json:"is_suspended"`
}

func (h *Handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := h.DB.Query(
		"SELECT id, name, email, role, is_suspended FROM users ORDER BY created_at DESC",
	)
	if err != nil {
		http.Error(w, "Eroare server", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var u User
		var name sql.NullString
		if err := rows.Scan(&u.ID, &name, &u.Email, &u.Role, &u.IsSuspended); err != nil {
			continue
		}
		u.Name = name.String
		users = append(users, u)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (h *Handler) SuspendUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var isSuspended bool
	err := h.DB.QueryRow("SELECT is_suspended FROM users WHERE id = $1", id).Scan(&isSuspended)
	if err != nil {
		http.Error(w, "User negasit", http.StatusNotFound)
		return
	}

	_, err = h.DB.Exec("UPDATE users SET is_suspended = $1 WHERE id = $2", !isSuspended, id)
	if err != nil {
		http.Error(w, "Eroare server", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"suspended": !isSuspended})
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	_, err := h.DB.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		http.Error(w, "Eroare server", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
