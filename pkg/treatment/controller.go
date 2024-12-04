package treatment

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"vet-clinic-api/config"
	"vet-clinic-api/database/dbmodel"
	"vet-clinic-api/pkg/models"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func CreateTreatmentHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.TreatmentCreateRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		date, err := time.Parse(time.RFC3339, req.Date)
		if err != nil {
			http.Error(w, "Invalid date format", http.StatusBadRequest)
			return
		}

		treatment := dbmodel.Treatment{
			VisitID:     req.VisitID,
			Name:        req.Name,
			Description: req.Description,
			Date:        date,
		}
		if err := cfg.TreatmentRepository.Create(&treatment); err != nil {
			http.Error(w, "Error creating treatment", http.StatusInternalServerError)
			return
		}
		response := models.TreatmentResponse{
			ID:          treatment.ID,
			VisitID:     treatment.VisitID,
			Name:        treatment.Name,
			Description: treatment.Description,
			Date:        treatment.Date.Format(time.RFC3339),
		}

		render.JSON(w, r, response)
	}
}

func GetAllTreatmentsHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		treatments, err := cfg.TreatmentRepository.FindAll()
		if err != nil {
			http.Error(w, "Error fetching treatments", http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, treatments)
	}
}

func GetTreatmentByIDHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		treatment, err := cfg.TreatmentRepository.FindByID(uint(id))
		if err != nil {
			http.Error(w, "treatment not found", http.StatusNotFound)
			return
		}

		render.JSON(w, r, treatment)
	}
}

func UpdateTreatmentHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		var updatedTreatment dbmodel.Treatment
		if err := json.NewDecoder(r.Body).Decode(&updatedTreatment); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		updatedTreatment.ID = uint(id)
		if err := cfg.TreatmentRepository.Update(&updatedTreatment); err != nil {
			http.Error(w, "Error updating treatment", http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, updatedTreatment)
	}
}

func DeleteTreatmentHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		if err := cfg.TreatmentRepository.Delete(uint(id)); err != nil {
			http.Error(w, "Error deleting treatment", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
