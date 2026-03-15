package support

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"clarity-gym/internal/auth"
)

type Handler struct {
	DB *sql.DB
}

type SupportRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Category string `json:"category"`
	Message  string `json:"message"`
}

func (h *Handler) SendRequest(w http.ResponseWriter, r *http.Request) {
	userID, _ := r.Context().Value(auth.UserIDKey).(string)

	var req SupportRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Date invalide", http.StatusBadRequest)
		return
	}

	if req.Name == "" || req.Email == "" || req.Category == "" || req.Message == "" {
		http.Error(w, "Toate campurile sunt obligatorii", http.StatusBadRequest)
		return
	}

	var id string
	err := h.DB.QueryRow(
		"INSERT INTO support_requests (user_id, name, email, category, message) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		userID, req.Name, req.Email, req.Category, req.Message,
	).Scan(&id)

	if err != nil {
		http.Error(w, "Eroare server", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": id, "status": "ok"})
}

func (h *Handler) GetAllRequests(w http.ResponseWriter, r *http.Request) {
	rows, err := h.DB.Query(`
		SELECT sr.id, sr.name, sr.email, sr.category, sr.message, sr.status, sr.created_at, u.role
		FROM support_requests sr
		JOIN users u ON u.id = sr.user_id
		ORDER BY sr.created_at DESC
	`)

	if err != nil {
		http.Error(w, "Eroare server", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type Request struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		Email     string `json:"email"`
		Category  string `json:"category"`
		Message   string `json:"message"`
		Status    string `json:"status"`
		CreatedAt string `json:"created_at"`
		Role      string `json:"role"`
	}

	requests := []Request{}
	for rows.Next() {
		var req Request
		if err := rows.Scan(&req.ID, &req.Name, &req.Email, &req.Category, &req.Message, &req.Status, &req.CreatedAt, &req.Role); err != nil {
			continue
		}
		requests = append(requests, req)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(requests)
}

func (h *Handler) CloseRequest(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID lipsa", http.StatusBadRequest)
		return
	}

	_, err := h.DB.Exec("UPDATE support_requests SET status = 'closed' WHERE id = $1", id)
	if err != nil {
		http.Error(w, "Eroare server", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "closed"})
}
