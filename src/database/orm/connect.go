package orm

import (
	"errors"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Koneksi ke Database dengan GORM
// var godotenv helpers.Godot

func New() (*gorm.DB, error) {

	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	password := os.Getenv("PASS")
	dbName := os.Getenv("DB")

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s", host, user, password, dbName)

	gormDb, err := gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		return nil, errors.New("failed connection to database")
	}

	db, err := gormDb.DB()
	if err != nil {
		return nil, errors.New("failed connection to database")
	}

	db.SetConnMaxIdleTime(100)
	db.SetMaxOpenConns(10)

	return gormDb, nil
}
