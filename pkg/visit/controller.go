package visit

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

func CreateVisitHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.VisitCreateRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		date, err := time.Parse(time.RFC3339, req.Date)
		if err != nil {
			http.Error(w, "Invalid date format", http.StatusBadRequest)
			return
		}

		visit := dbmodel.Visit{
			CatID:        req.CatID,
			Date:         date,
			Reason:       req.Reason,
			Veterinarian: req.Veterinarian,
		}

		if err := cfg.VisitRepository.Create(&visit); err != nil {
			http.Error(w, "Error creating visit", http.StatusInternalServerError)
			return
		}

		response := models.VisitResponse{
			ID:           visit.ID,
			Date:         visit.Date.Format(time.RFC3339),
			Reason:       visit.Reason,
			Veterinarian: visit.Veterinarian,
		}

		render.JSON(w, r, response)
	}
}

func GetAllVisitsHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		visits, err := cfg.VisitRepository.FindAll()
		if err != nil {
			http.Error(w, "Error fetching visits", http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, visits)
	}
}

func GetVisitByIDHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		visit, err := cfg.VisitRepository.FindByID(uint(id))
		if err != nil {
			http.Error(w, "Visit not found", http.StatusNotFound)
			return
		}

		render.JSON(w, r, visit)
	}
}

func UpdateVisitHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		var updatedVisit dbmodel.Visit
		if err := json.NewDecoder(r.Body).Decode(&updatedVisit); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		updatedVisit.ID = uint(id)
		if err := cfg.VisitRepository.Update(&updatedVisit); err != nil {
			http.Error(w, "Error updating Visit", http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, updatedVisit)
	}
}

func DeleteVisitHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		if err := cfg.VisitRepository.Delete(uint(id)); err != nil {
			http.Error(w, "Error deleting Visit", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func FilterVisitsHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var visits []dbmodel.Visit

		date := r.URL.Query().Get("date")
		veterinarian := r.URL.Query().Get("veterinarian")
		reason := r.URL.Query().Get("reason")

		query := cfg.DB.Model(&dbmodel.Visit{})

		if date != "" {
			query = query.Where("date = ?", date)
		}
		if veterinarian != "" {
			query = query.Where("veterinarian = ?", veterinarian)
		}
		if reason != "" {
			query = query.Where("reason LIKE ?", "%"+reason+"%")
		}

		if err := query.Find(&visits).Error; err != nil {
			http.Error(w, "Error fetching visits", http.StatusInternalServerError)
			return
		}

		response := make([]models.VisitResponse, len(visits))
		for i, visit := range visits {
			response[i] = models.VisitResponse{
				ID:           visit.ID,
				Date:         visit.Date.Format(time.RFC3339),
				Reason:       visit.Reason,
				Veterinarian: visit.Veterinarian,
			}
		}

		render.JSON(w, r, response)
	}
}
