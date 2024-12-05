package database

import (
    "github.com/glebarez/sqlite"
    "gorm.io/gorm"
    "log"
	"animal-api/database/dbmodel"
	"gorm.io/gorm"

)

var DB *gorm.DB

func InitDatabase() {
    var err error
    DB, err = gorm.Open(sqlite.Open("animal_api.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
	DB.AutoMigrate(&dbmodel.AgeEntry{})
    log.Println("Database connected")
}

func Migrate(db *gorm.DB) {
	log.Println("Running migrations...")
	dbmodel.MigrateCat(db)
	dbmodel.MigrateVisit(db)
	dbmodel.MigrateTreatment(db)
	dbmodel.MigrateUser(db)
	log.Println("Migrations complete.")
}
