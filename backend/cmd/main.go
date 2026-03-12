package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"clarity-gym/config"
	"clarity-gym/internal/auth"
	"clarity-gym/internal/user"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
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

	userHandler := &user.Handler{
		DB: db,
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Clarity Gym API"))
	})

	// Auth routes
	r.Post("/api/auth/register", authHandler.Register)
	r.Post("/api/auth/login", authHandler.Login)

	// Admin routes - protejate cu JWT
	r.Group(func(r chi.Router) {
		r.Use(auth.AuthMiddleware(cfg.JWTSecret))
		r.Get("/api/admin/users", userHandler.GetAllUsers)
		r.Put("/api/admin/users/{id}/suspend", userHandler.SuspendUser)
		r.Delete("/api/admin/users/{id}", userHandler.DeleteUser)
	})

	fmt.Println("🚀 Server pornit pe portul", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, c.Handler(r)))
}
