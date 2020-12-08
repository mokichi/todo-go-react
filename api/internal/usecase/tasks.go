package usecase

import (
	"gorm.io/gorm"
	"todo/internal/entity"
)

type TasksUsecase struct {
	DB *gorm.DB
}

func (u TasksUsecase) FindByCompleted(completed bool) []entity.Task {
	tasks := make([]entity.Task, 0)
	u.DB.Where("completed = ?", completed).Order("updated_at DESC").Find(&tasks)
	return tasks
}

func (u TasksUsecase) First(id interface{}) (entity.Task, error) {
	task := entity.Task{}
	err := u.DB.First(&task, id).Error
	return task, err
}

func (u TasksUsecase) Create(task *entity.Task) error {
	return u.DB.Create(&task).Error
}

func (u TasksUsecase) Update(task *entity.Task) error {
	return u.DB.Save(&task).Error
}

func (u TasksUsecase) Delete(task *entity.Task) error {
	return u.DB.Delete(&task).Error
}
