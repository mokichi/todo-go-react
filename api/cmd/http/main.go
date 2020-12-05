package main

import (
	"fmt"
	"todo/internal/entity"
	"todo/internal/orm"
)

func main() {
	db, err := orm.NewDB()
	if err != nil {
		panic(err)
	}

	task := entity.Task{Title: "Test"}
	err = db.Create(&task).Error
	fmt.Println(task, err)

	err = db.Model(&task).Updates(map[string]interface{}{"Completed": true}).Error
	fmt.Println(task, err)

	err = db.Delete(&task).Error
	fmt.Println(err)

	t := entity.Task{}
	err = db.First(&t, task.ID).Error
	fmt.Println(t, err)
}
