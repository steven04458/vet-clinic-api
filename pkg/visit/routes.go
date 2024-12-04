package visit

import (
	"vet-clinic-api/config"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router, cfg *config.Config) {
	r.Route("/visits", func(r chi.Router) {
		r.Post("/", CreateVisitHandler(cfg))
		r.Get("/", GetAllVisitsHandler(cfg))
		r.Get("/filter", FilterVisitsHandler(cfg))
		r.Put("/{id}", UpdateVisitHandler(cfg))
		r.Delete("/{id}", DeleteVisitHandler(cfg))
	})
}
