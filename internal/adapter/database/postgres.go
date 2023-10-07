package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDB struct {
	db *gorm.DB
}

func NewPostgresDB() (*gorm.DB, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	strConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)

	gormDB, err := gorm.Open(postgres.Open(strConn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	postgres := postgresDB{db: gormDB}

	return postgres.db, nil

}
