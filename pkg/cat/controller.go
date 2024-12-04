package cat

import (
	"encoding/json"
	"net/http"
	"strconv"
	"vet-clinic-api/config"
	"vet-clinic-api/database/dbmodel"
	"vet-clinic-api/pkg/models"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func CreateCatHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.CatCreateRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		cat := dbmodel.Cat{
			Name:   req.Name,
			Age:    req.Age,
			Breed:  req.Breed,
			Weight: req.Weight,
		}

		if err := cfg.CatRepository.Create(&cat); err != nil {
			http.Error(w, "Error creating cat", http.StatusInternalServerError)
			return
		}

		response := models.CatResponse{
			ID:     cat.ID,
			Name:   cat.Name,
			Age:    cat.Age,
			Breed:  cat.Breed,
			Weight: cat.Weight,
		}

		render.JSON(w, r, response)
	}
}

func GetAllCatsHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cats, err := cfg.CatRepository.FindAll()
		if err != nil {
			http.Error(w, "Error fetching cats", http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, cats)
	}
}

func GetCatByIDHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		cat, err := cfg.CatRepository.FindByID(uint(id))
		if err != nil {
			http.Error(w, "Cat not found", http.StatusNotFound)
			return
		}

		render.JSON(w, r, cat)
	}
}

func UpdateCatHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		var updatedCat dbmodel.Cat
		if err := json.NewDecoder(r.Body).Decode(&updatedCat); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		updatedCat.ID = uint(id)
		if err := cfg.CatRepository.Update(&updatedCat); err != nil {
			http.Error(w, "Error updating cat", http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, updatedCat)
	}
}

func DeleteCatHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		if err := cfg.CatRepository.Delete(uint(id)); err != nil {
			http.Error(w, "Error deleting cat", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func HistoryCatHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		history, err := cfg.CatRepository.FindHistoryWithTreatments(uint(id)) // Updated to include treatments
		if err != nil {
			http.Error(w, "Error fetching history", http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, history)
	}
}
