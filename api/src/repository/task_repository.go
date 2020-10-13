package repository

import (
	"gorm.io/gorm"
	"entity"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository() TaskRepository {
	return TaskRepository{db: db}
}

func (repo *TaskRepository) Find(id uint) (entity.Task, error) {
	task := entity.Task{}
	err := repo.db.First(&task, id).Error
	return task, err
}

func (repo *TaskRepository) Create(task *entity.Task) error {
	return repo.db.Create(task).Error
}

func (repo *TaskRepository) Update(task *entity.Task, attrs map[string]interface{}) error {
	return repo.db.Model(task).Updates(attrs).Error
}

func (repo *TaskRepository) Delete(task *entity.Task) error {
	return repo.db.Delete(task).Error
}

func init() {
	db.AutoMigrate(&entity.Task{})
}
