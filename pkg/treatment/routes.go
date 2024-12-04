package treatment

import (
	"vet-clinic-api/config"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router, cfg *config.Config) {
	r.Route("/treatments", func(r chi.Router) {
		r.Post("/", CreateTreatmentHandler(cfg))
		r.Get("/", GetAllTreatmentsHandler(cfg))
		r.Get("/{id}", GetTreatmentByIDHandler(cfg))
		r.Put("/{id}", UpdateTreatmentHandler(cfg))
		r.Delete("/{id}", DeleteTreatmentHandler(cfg))
	})
}
