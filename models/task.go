package models

import "time"

type Task struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string    `gorm:"not null" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	Status      string    `gorm:"type:text; not null" json:"status" default:"todo"`
	Priority    string    `gorm:"type:text; not null" json:"priority" default:"low"`
	DueDate     time.Time `json:"dueDate"`
	Subtasks    []Subtask `json:"-"`
	Labels      []Label   `gorm:"many2many:task_labels;" json:"-"`
	Comments    []Comment `json:"comments"`
	ProjectID   uint      `gorm:"not null" json:"projectId"`
	Project     Project   `gorm:"foreignKey:ProjectID" json:"-"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
type TaskFilter struct {
	Title    string `json:"title"`
	Status   string `json:"status"`
	Label    string `json:"label"`
	Priority string `json:"priority"`
}
