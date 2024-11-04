package models

import "time"

type Task struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string    `gorm:"not null" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	Status      string    `gorm:"type:text; not null" json:"status" default:"todo"`
	Priority    string    `gorm:"type:text; not null" json:"priority" default:"low"`
	DueDate     time.Time `json:"dueDate"`
	UserID      uint      `gorm:"not null" json:"userId"`
	User        User      `gorm:"foreignKey:UserID" json:"-"`
	Labels      []Label   `gorm:"many2many:task_labels;" json:"labels"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
