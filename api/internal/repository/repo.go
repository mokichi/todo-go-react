package repository

import "gorm.io/gorm"

type Repo interface {
	Find(entity interface{}, id uint) error
	Create(entity interface{}) error
	Update(entity interface{}, attrs interface{}) error
	Delete(entity interface{}) error
}

type repo struct {
	db *gorm.DB
}

func NewRepo() Repo {
	return repo{db: db}
}

func (r repo) Find(entity interface{}, id uint) error {
	return r.db.First(entity, id).Error
}

func (r repo) Create(entity interface{}) error {
	return r.db.Create(entity).Error
}

func (r repo) Update(entity interface{}, attrs interface{}) error {
	return r.db.Model(entity).Updates(attrs).Error
}

func (r repo) Delete(entity interface{}) error {
	return r.db.Delete(entity).Error
}
