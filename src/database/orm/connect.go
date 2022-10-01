package orm

import (
	"errors"
	"fmt"

	"github.com/adiet95/gorent-api/src/helpers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Koneksi ke Database dengan GORM
// var godotenv helpers.Godot

func New() (*gorm.DB, error) {

	host := helpers.Godotenv("HOST")
	user := helpers.Godotenv("USER")
	password := helpers.Godotenv("PASS")
	dbName := helpers.Godotenv("DB")

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
