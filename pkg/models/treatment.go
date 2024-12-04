package models

type TreatmentResponse struct {
	ID          uint   `json:"id"`
	VisitID     uint   `json:"visit_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Date        string `json:"date"`
}

type TreatmentCreateRequest struct {
	VisitID     uint   `json:"visit_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Date        string `json:"date" validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
}
