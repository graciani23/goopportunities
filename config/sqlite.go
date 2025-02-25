package config

import (
	"os"

	"github.com/graciani23/goopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")

	dbPath := "./db/main.db"

	// Check if the database file exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Infof("Database file does not exist, creating a new one")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			logger.Errf("Error creating database directory: %v", err)
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			logger.Errf("Error creating database file: %v", err)
			return nil, err
		}
		file.Close()
	}

	// Open a database connection
	db, error := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if error != nil {
		logger.Errf("Error initializing sqlite: %v", error)
		return nil, error
	}

	// Migrate the schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errf("Error migrating schema: %v", err)
		return nil, err
	}
	return db, nil
}
