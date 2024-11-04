package models

import "time"

type Subtask struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title     string    `gorm:"not null" json:"title"`
	Status    string    `gorm:"type:text; not null" json:"status" default:"pending"`
	TaskID    uint      `gorm:"not null" json:"taskId"`
	Task      Task      `gorm:"foreignKey:TaskID" json:"-"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
