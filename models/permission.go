package models

type Permission struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"unique;not null" json:"name"` // e.g., "create_task"
	Roles []Role `gorm:"many2many:role_permissions;" json:"-"`
}
