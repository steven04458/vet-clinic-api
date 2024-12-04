package database

import (
	"log"
	"vet-clinic-api/database/dbmodel"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() error {
	var err error
	DB, err = gorm.Open(sqlite.Open("vetclinic.db"), &gorm.Config{})
	if err != nil {
		log.Printf("Erreur de connexion à la base de données : %v", err)
		return err
	}
	log.Println("Connexion à la base de données réussie.")
	return nil
}

func Migrate() error {
	err := DB.AutoMigrate(&dbmodel.Cat{}, &dbmodel.Visit{}, &dbmodel.Treatment{})
	if err != nil {
		log.Printf("Erreur lors de la migration : %v", err)
		return err
	}
	log.Println("Migration des modèles terminée.")
	return nil
}
