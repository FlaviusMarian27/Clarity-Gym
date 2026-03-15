package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"clarity-gym/config"
	"clarity-gym/internal/auth"
	"clarity-gym/internal/collaboration"
	"clarity-gym/internal/trainer"
	"clarity-gym/internal/user"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
	"github.com/rs/cors"

	"clarity-gym/internal/support"

	"clarity-gym/internal/subscription"
)

func main() {
	cfg := config.Load()

	db, err := sql.Open("postgres", cfg.DBConnectionString())
	if err != nil {
		log.Fatal("Nu ma pot conecta la baza de date:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Baza de date nu raspunde:", err)
	}

	fmt.Println("✅ Conectat la PostgreSQL!")

	authHandler := &auth.Handler{
		DB:        db,
		JWTSecret: cfg.JWTSecret,
	}
	userHandler := &user.Handler{DB: db}
	trainerHandler := &trainer.Handler{DB: db}
	collaborationHandler := &collaboration.Handler{DB: db}
	supportHandler := &support.Handler{DB: db}
	subscriptionHandler := &subscription.Handler{DB: db}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	// Servire fisiere statice
	r.Handle("/uploads/*", http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads"))))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Clarity Gym API"))
	})

	// Auth routes
	r.Post("/api/auth/register", authHandler.Register)
	r.Post("/api/auth/login", authHandler.Login)

	// Public routes
	r.Get("/api/trainers", trainerHandler.GetAllTrainers)

	// Planuri publice
	r.Get("/api/plans", subscriptionHandler.GetAllPlans)

	// Protected routes
	r.Group(func(r chi.Router) {
		r.Use(auth.AuthMiddleware(cfg.JWTSecret))

		// Admin
		r.Get("/api/admin/users", userHandler.GetAllUsers)
		r.Put("/api/admin/users/{id}/suspend", userHandler.SuspendUser)
		r.Delete("/api/admin/users/{id}", userHandler.DeleteUser)

		// Colaborare
		r.Post("/api/collaborations", collaborationHandler.SendRequest)
		r.Get("/api/collaborations/my", collaborationHandler.GetMyRequests)
		r.Get("/api/collaborations/clients", collaborationHandler.GetMyClients)
		r.Put("/api/collaborations/{id}/respond", collaborationHandler.RespondToRequest)
		r.Get("/api/collaborations/status", collaborationHandler.GetMyStatus)
		r.Post("/api/collaborations/seen", collaborationHandler.MarkAsSeen)

		// Profil
		r.Get("/api/profile", userHandler.GetProfile)
		r.Put("/api/profile", userHandler.UpdateProfile)
		r.Post("/api/profile/avatar", userHandler.UploadAvatar)

		// Suport
		r.Post("/api/support", supportHandler.SendRequest)
		r.Get("/api/support", supportHandler.GetAllRequests)
		r.Put("/api/support/close", supportHandler.CloseRequest)

		// In grupul protejat
		r.Post("/api/plans", subscriptionHandler.AddPlan)
		r.Put("/api/plans/{id}", subscriptionHandler.UpdatePlan)
		r.Delete("/api/plans/{id}", subscriptionHandler.DeletePlan)
	})

	fmt.Println("🚀 Server pornit pe portul", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, c.Handler(r)))
}
