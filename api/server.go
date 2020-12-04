package main

import (
	"fmt"
	"todo/repository"
	"todo/entity"
)

func main() {
	repo := repository.NewTaskRepository()
	task := entity.Task{Title: "Test"}
	err := repo.Create(&task)
	fmt.Println(task, err)

	err = repo.Update(&task, map[string]interface{}{"Completed": true})
	fmt.Println(task, err)

	err = repo.Delete(&task)
	fmt.Println(err)

	t, err := repo.Find(task.ID)
	fmt.Println(t, err)
}
