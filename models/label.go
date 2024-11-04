package models

type Label struct {
	ID    uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name  string `gorm:"type:text; not null" json:"name" default:"work"`
	Tasks []Task `gorm:"many2many:task_labels;" json:"-"`
}
