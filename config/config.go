package config

import (
	"log"
	"vet-clinic-api/database/dbmodel"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	DB                  *gorm.DB
	CatRepository       dbmodel.CatRepository
	VisitRepository     dbmodel.VisitRepository
	TreatmentRepository dbmodel.TreatmentRepository
}

func New() (*Config, error) {
	db, err := gorm.Open(sqlite.Open("vetclinic.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Impossible de se connecter à la base de données : %v", err)
		return nil, err
	}

	if err := db.AutoMigrate(&dbmodel.Cat{}, &dbmodel.Visit{}, &dbmodel.Treatment{}); err != nil {
		log.Fatalf("Erreur lors de la migration : %v", err)
		return nil, err
	}

	catRepo := dbmodel.NewCatRepository(db)
	visitRepo := dbmodel.NewVisitRepository(db)
	treatmentRepo := dbmodel.NewTreatmentRepository(db)

	config := &Config{
		DB:                  db,
		CatRepository:       catRepo,
		VisitRepository:     visitRepo,
		TreatmentRepository: treatmentRepo,
	}

	return config, nil
}
