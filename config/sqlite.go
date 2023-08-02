package config

import (
	"os"

	"github.com/CarlosGenuino/go-opportunities/schema"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitilizeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	//create
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("Database file not exists, creating...")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			logger.Errorf("Error opening conection with Sqlite. %v", err)
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			logger.Errorf("Error creating file sqlite. %v", err)
			return nil, err
		}
		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("Error on establish database conection. %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schema.Opening{})
	if err != nil {
		logger.Errorf("Error on migrate schema. %v", err)
		return nil, err
	}

	return db, nil

}
