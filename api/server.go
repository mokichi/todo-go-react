package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title     string
	Completed bool `gorm:"default:false"`
}

func main() {
	dsn := "host=db user=postgres password=postgres dbname=todo_dev port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Task{})

	task := Task{Title: "Test"}
	db.Create(&task)
	fmt.Println(task)

	db.Model(&task).Updates(map[string]interface{}{"Completed": true})
	fmt.Println(task)

	db.Delete(&task)
}
