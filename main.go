package main

import (
	"log"
	"net/http"
	"vet-clinic-api/config"
	"vet-clinic-api/pkg/authentication"
	"vet-clinic-api/pkg/cat"

	"github.com/go-chi/chi/v5"
)

func main() {
	cfg := config.NewConfig()

	// Migrate models
	cfg.DB.AutoMigrate(&dbmodel.Cat{}, &dbmodel.User{})

	// Initialize controllers
	cat.InitController(cfg)
	authentication.InitAuthController(cfg)

	// Setup router
	r := chi.NewRouter()

	// Routes
	r.Group(func(r chi.Router) {
		r.Use(authentication.JWTMiddleware)
		cat.RegisterRoutes(r)
	})

	r.Post("/login", authentication.Login)

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
