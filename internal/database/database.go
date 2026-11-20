package database

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"portfolio/internal/models"
)

// Connect opens a GORM connection to Postgres and applies AutoMigrate for
// all domain models. AutoMigrate is safe to run on every boot (it only
// creates/alters, never drops), which keeps local dev and small deployments
// simple. See /migrations for the raw-SQL equivalents used by golang-migrate
// in environments that want explicit, reviewable migration history.
func Connect(dsn string) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	for attempt := 1; attempt <= 5; attempt++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Warn),
		})
		if err == nil {
			break
		}
		log.Printf("database: connection attempt %d/5 failed: %v", attempt, err)
		time.Sleep(time.Duration(attempt) * time.Second)
	}
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(
		&models.Project{},
		&models.Certification{},
		&models.Article{},
		&models.ContactMessage{},
	); err != nil {
		return nil, err
	}

	return db, nil
}
