package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlDB() (*gorm.DB, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	dbHost := os.Getenv("DBHOST")
	dbPort := os.Getenv("DBPORT")
	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASSWORD")
	dbName := os.Getenv("DBNAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return gormDB, nil

}
