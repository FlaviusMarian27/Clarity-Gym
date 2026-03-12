package collaboration

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"clarity-gym/internal/auth"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	DB *sql.DB
}

type CollaborationRequest struct {
	TrainerID string `json:"trainer_id"`
}

type Collaboration struct {
	ID        string `json:"id"`
	ClientID  string `json:"client_id"`
	TrainerID string `json:"trainer_id"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}

func (h *Handler) SendRequest(w http.ResponseWriter, r *http.Request) {
	clientID, _ := r.Context().Value(auth.UserIDKey).(string)

	var req CollaborationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Date invalide", http.StatusBadRequest)
		return
	}

	var existingID string
	err := h.DB.QueryRow(
		"SELECT id FROM collaboration_requests WHERE client_id = $1 AND trainer_id = $2 AND status = 'pending'",
		clientID, req.TrainerID,
	).Scan(&existingID)

	if err == nil {
		http.Error(w, "Ai deja o cerere trimisă acestui antrenor", http.StatusConflict)
		return
	}

	var id string
	err = h.DB.QueryRow(
		"INSERT INTO collaboration_requests (client_id, trainer_id) VALUES ($1, $2) RETURNING id",
		clientID, req.TrainerID,
	).Scan(&id)

	if err != nil {
		http.Error(w, "Eroare server", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": id, "status": "pending"})
}

func (h *Handler) GetMyRequests(w http.ResponseWriter, r *http.Request) {
	trainerID, _ := r.Context().Value(auth.UserIDKey).(string)

	rows, err := h.DB.Query(`
		SELECT cr.id, cr.client_id, u.name, cr.status, cr.created_at
		FROM collaboration_requests cr
		JOIN users u ON u.id = cr.client_id
		WHERE cr.trainer_id = $1
		ORDER BY cr.created_at DESC
	`, trainerID)

	if err != nil {
		http.Error(w, "Eroare server", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type Request struct {
		ID         string `json:"id"`
		ClientID   string `json:"client_id"`
		ClientName string `json:"client_name"`
		Status     string `json:"status"`
		CreatedAt  string `json:"created_at"`
	}

	requests := []Request{}
	for rows.Next() {
		var req Request
		if err := rows.Scan(&req.ID, &req.ClientID, &req.ClientName, &req.Status, &req.CreatedAt); err != nil {
			continue
		}
		requests = append(requests, req)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(requests)
}

func (h *Handler) RespondToRequest(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	type Response struct {
		Status string `json:"status"`
	}

	var resp Response
	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		http.Error(w, "Date invalide", http.StatusBadRequest)
		return
	}

	if resp.Status != "accepted" && resp.Status != "rejected" {
		http.Error(w, "Status invalid", http.StatusBadRequest)
		return
	}

	_, err := h.DB.Exec(
		"UPDATE collaboration_requests SET status = $1 WHERE id = $2",
		resp.Status, id,
	)

	if err != nil {
		http.Error(w, "Eroare server", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": resp.Status})
}

func (h *Handler) GetMyClients(w http.ResponseWriter, r *http.Request) {
	trainerID, _ := r.Context().Value(auth.UserIDKey).(string)

	rows, err := h.DB.Query(`
		SELECT u.id, u.name, u.email, cr.created_at
		FROM collaboration_requests cr
		JOIN users u ON u.id = cr.client_id
		WHERE cr.trainer_id = $1 AND cr.status = 'accepted'
		ORDER BY cr.created_at DESC
	`, trainerID)

	if err != nil {
		http.Error(w, "Eroare server", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type Client struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
		Since string `json:"since"`
	}

	clients := []Client{}
	for rows.Next() {
		var c Client
		if err := rows.Scan(&c.ID, &c.Name, &c.Email, &c.Since); err != nil {
			continue
		}
		clients = append(clients, c)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clients)
}

func (h *Handler) GetMyStatus(w http.ResponseWriter, r *http.Request) {
	clientID, _ := r.Context().Value(auth.UserIDKey).(string)

	rows, err := h.DB.Query(`
		SELECT cr.id, cr.trainer_id, u.name, cr.status, cr.created_at
		FROM collaboration_requests cr
		JOIN users u ON u.id = cr.trainer_id
		WHERE cr.client_id = $1
		ORDER BY cr.created_at DESC
	`, clientID)

	if err != nil {
		http.Error(w, "Eroare server", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type Status struct {
		ID          string `json:"id"`
		TrainerID   string `json:"trainer_id"`
		TrainerName string `json:"trainer_name"`
		Status      string `json:"status"`
		CreatedAt   string `json:"created_at"`
	}

	statuses := []Status{}
	for rows.Next() {
		var s Status
		if err := rows.Scan(&s.ID, &s.TrainerID, &s.TrainerName, &s.Status, &s.CreatedAt); err != nil {
			continue
		}
		statuses = append(statuses, s)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(statuses)
}
