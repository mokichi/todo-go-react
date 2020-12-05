package orm

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"time"
)

func NewDB() (*gorm.DB, error) {
	dsn := os.Getenv("DATABASE_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}
