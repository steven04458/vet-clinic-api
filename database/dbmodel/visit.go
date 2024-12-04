package dbmodel

import (
	"time"

	"gorm.io/gorm"
)

type Visit struct {
	gorm.Model
	ID           uint        `gorm:"primaryKey"`
	CatID        uint        `gorm:"not null" json:"cat_id"`
	Date         time.Time   `gorm:"not null" json:"date"`
	Reason       string      `gorm:"size:255" json:"reason"`
	Veterinarian string      `gorm:"size:255" json:"veterinarian"`
	Treatments   []Treatment `gorm:"foreignKey:VisitID" json:"treatments"`
}

type VisitRepository interface {
	Create(visit *Visit) error
	FindByID(id uint) (*Visit, error)
	FindAll() ([]Visit, error)
	Update(visit *Visit) error
	Delete(id uint) error
}

type visitRepository struct {
	db *gorm.DB
}

func NewVisitRepository(db *gorm.DB) VisitRepository {
	return &visitRepository{db: db}
}

func (r *visitRepository) Create(visit *Visit) error {
	return r.db.Create(visit).Error
}

func (r *visitRepository) FindByID(id uint) (*Visit, error) {
	var visit Visit
	err := r.db.First(&visit, id).Error
	return &visit, err
}

func (r *visitRepository) FindAll() ([]Visit, error) {
	var visits []Visit
	err := r.db.Find(&visits).Error
	return visits, err
}

func (r *visitRepository) Update(visit *Visit) error {
	return r.db.Save(visit).Error
}

func (r *visitRepository) Delete(id uint) error {
	return r.db.Delete(&Visit{}, id).Error
}

func (r *visitRepository) FindByIDWithTreatments(id uint) (*Visit, error) {
	var visit Visit
	err := r.db.Preload("Treatments").First(&visit, id).Error // Preload treatments
	return &visit, err
}
