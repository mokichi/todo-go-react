package entity

import "time"

type Task struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `gorm:"default:false" json:"completed"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
