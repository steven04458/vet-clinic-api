package models

type VisitResponse struct {
	ID           uint   `json:"id"`
	Date         string `json:"date"`
	Reason       string `json:"reason"`
	Veterinarian string `json:"veterinarian"`
}

type VisitCreateRequest struct {
	CatID        uint   `json:"cat_id" validate:"required"`
	Date         string `json:"date" validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
	Reason       string `json:"reason" validate:"required"`
	Veterinarian string `json:"veterinarian" validate:"required"`
}
