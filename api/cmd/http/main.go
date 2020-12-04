package main

import (
	"fmt"
	"todo/internal/entity"
	"todo/internal/repository"
)

func main() {
	repo := repository.NewRepo()
	task := entity.Task{Title: "Test"}
	err := repo.Create(&task)
	fmt.Println(task, err)

	err = repo.Update(&task, map[string]interface{}{"Completed": true})
	fmt.Println(task, err)

	err = repo.Delete(&task)
	fmt.Println(err)

	t := entity.Task{}
	err = repo.Find(&t, task.ID)
	fmt.Println(t, err)
}
