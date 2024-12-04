package dbmodel

import "gorm.io/gorm"

type Cat struct {
	gorm.Model
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Breed  string  `json:"breed"`
	Weight float32 `json:"weight"`
	Visits []Visit `gorm:"foreignKey:CatID" json:"visits"` // Relation avec les visites
}

type CatRepository interface {
	Create(cat *Cat) error
	FindByID(id uint) (*Cat, error)
	FindAll() ([]Cat, error)
	Update(cat *Cat) error
	Delete(id uint) error
	FindHistory(id uint) ([]Visit, error)
	FindHistoryWithTreatments(catID uint) ([]Visit, error) // New method
}

type catRepository struct {
	db *gorm.DB
}

func NewCatRepository(db *gorm.DB) CatRepository {
	return &catRepository{db: db}
}

func (r *catRepository) Create(cat *Cat) error {
	return r.db.Create(cat).Error
}

func (r *catRepository) FindByID(id uint) (*Cat, error) {
	var cat Cat
	err := r.db.First(&cat, id).Error
	return &cat, err
}

func (r *catRepository) FindAll() ([]Cat, error) {
	var cats []Cat
	err := r.db.Find(&cats).Error
	return cats, err
}

func (r *catRepository) Update(cat *Cat) error {
	return r.db.Save(cat).Error
}

func (r *catRepository) Delete(id uint) error {
	return r.db.Delete(&Cat{}, id).Error
}

func (r *catRepository) FindHistory(id uint) ([]Visit, error) {
	var visits []Visit
	err := r.db.Model(&Cat{}).Where("id = ?", id).Association("Visits").Find(&visits)
	return visits, err
}

func (r *catRepository) FindHistoryWithTreatments(catID uint) ([]Visit, error) {
	var visits []Visit
	err := r.db.Preload("Treatments").Where("cat_id = ?", catID).Find(&visits).Error // Preload treatments
	return visits, err
}
