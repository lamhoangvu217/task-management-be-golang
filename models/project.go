package models

import (
	"time"
)

type Project struct {
	ID          uint              `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string            `gorm:"not null" json:"title"`
	Description string            `gorm:"type:text" json:"description"`
	Status      string            `gorm:"type:varchar(20);default:'active'" json:"status"`
	StartDate   time.Time         `json:"start_date"`
	EndDate     time.Time         `json:"end_date"`
	OwnerID     uint              `gorm:"not null" json:"ownerId"`
	Owner       User              `gorm:"foreignKey:OwnerID" json:"-"`
	Users       []UserProjectRole `gorm:"foreignKey:ProjectID" json:"collaborators"`
	Tasks       []Task            `json:"-"`
	CreatedAt   time.Time         `json:"createdAt"`
	UpdatedAt   time.Time         `json:"updatedAt"`
}

type CollaboratorResponse struct {
	UserID   uint   `json:"userId"`
	Email    string `json:"email"`
	FullName string `json:"fullName"`
	Role     string `json:"role"`
}
