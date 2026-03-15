package subscription

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	DB *sql.DB
}

type Plan struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

func (h *Handler) GetAllPlans(w http.ResponseWriter, r *http.Request) {
	rows, err := h.DB.Query(`
		SELECT id, name, type, price, COALESCE(description, '')
		FROM subscription_plans
		ORDER BY price ASC
	`)
	if err != nil {
		http.Error(w, "Eroare server", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	plans := []Plan{}
	for rows.Next() {
		var p Plan
		if err := rows.Scan(&p.ID, &p.Name, &p.Type, &p.Price, &p.Description); err != nil {
			continue
		}
		plans = append(plans, p)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(plans)
}

func (h *Handler) UpdatePlan(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var plan Plan
	if err := json.NewDecoder(r.Body).Decode(&plan); err != nil {
		http.Error(w, "Date invalide", http.StatusBadRequest)
		return
	}

	_, err := h.DB.Exec(`
		UPDATE subscription_plans SET name = $1, type = $2, price = $3, description = $4
		WHERE id = $5
	`, plan.Name, plan.Type, plan.Price, plan.Description, id)

	if err != nil {
		http.Error(w, "Eroare server", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func (h *Handler) DeletePlan(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	_, err := h.DB.Exec("DELETE FROM subscription_plans WHERE id = $1", id)
	if err != nil {
		http.Error(w, "Eroare server", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func (h *Handler) AddPlan(w http.ResponseWriter, r *http.Request) {
	var plan Plan
	if err := json.NewDecoder(r.Body).Decode(&plan); err != nil {
		http.Error(w, "Date invalide", http.StatusBadRequest)
		return
	}

	var id string
	err := h.DB.QueryRow(`
		INSERT INTO subscription_plans (name, type, price, description)
		VALUES ($1, $2, $3, $4) RETURNING id
	`, plan.Name, plan.Type, plan.Price, plan.Description).Scan(&id)

	if err != nil {
		http.Error(w, "Eroare server", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": id, "status": "ok"})
}
