package cat

import (
	"vet-clinic-api/config"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router, cfg *config.Config) {
	r.Route("/cats", func(r chi.Router) {
		r.Post("/", CreateCatHandler(cfg))
		r.Get("/", GetAllCatsHandler(cfg))
		r.Get("/{id}", GetCatByIDHandler(cfg))
		r.Get("/{id}/history", HistoryCatHandler(cfg))
		r.Put("/{id}", UpdateCatHandler(cfg))
		r.Delete("/{id}", DeleteCatHandler(cfg))
	})
}
