// db_migrations.go
package main

import (
	"anime-kuring/models"

	"gorm.io/gorm"
)

// Migration represents a database migration
type Migration struct {
	ID      uint `gorm:"primaryKey"`
	Version string
}

// MigrateDB performs database migrations
func MigrateDB(db *gorm.DB, desiredVersion string) error {
	// Check if migration table exists
	if !db.Migrator().HasTable(&Migration{}) {
		// If not, create the migration table
		if err := db.AutoMigrate(&Migration{}, &models.User{}); err != nil {
			return err
		}

		// Insert the initial version
		if err := db.Create(&Migration{Version: desiredVersion}).Error; err != nil {
			return err
		}
	}

	// Retrieve the current version from the database
	var currentVersion Migration
	if err := db.First(&currentVersion).Error; err != nil {
		return err
	}

	// Compare the current version with the desired version
	if currentVersion.Version != desiredVersion {
		// Perform migrations if versions are different
		if err := db.AutoMigrate(&models.User{}); err != nil {
			return err
		}

		// Update the version in the database
		if err := db.Model(&currentVersion).Update("Version", desiredVersion).Error; err != nil {
			return err
		}
	}

	return nil
}
