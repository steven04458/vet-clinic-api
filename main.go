package main

import (
	"log"
	"net/http"
	"os"
	"vet-clinic-api/config"
	"vet-clinic-api/pkg/cat"
	"vet-clinic-api/pkg/treatment"
	"vet-clinic-api/pkg/visit"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Erreur lors de l'initialisation de la configuration : %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	render.SetContentType(render.ContentTypeJSON)

	r.Route("/api/v1", func(r chi.Router) {
		cat.RegisterRoutes(r, cfg)
		visit.RegisterRoutes(r, cfg)
		treatment.RegisterRoutes(r, cfg)
	})

	log.Printf("Serveur démarré sur le port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Erreur lors du démarrage du serveur : %v", err)
	}
}
