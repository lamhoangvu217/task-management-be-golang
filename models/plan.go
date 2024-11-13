package models

type Plan struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string `gorm:"type:text; not null" json:"name"`
	Description string `gorm:"type:text" json:"description"`
	Price       uint   `gorm:"not null" json:"price"`
}

type SubscribePlanRequest struct {
	PlanID uint `json:"planId" validate:"required"`
}
