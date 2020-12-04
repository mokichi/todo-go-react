package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"time"
)

var db *gorm.DB = newDB()

func newDB() *gorm.DB {
	dsn := os.Getenv("DATABASE_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}
