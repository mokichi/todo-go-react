package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"todo/internal/entity"
)

func main() {
	dsn := os.Getenv("DATABASE_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Task{})
}
