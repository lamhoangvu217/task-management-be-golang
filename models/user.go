package models

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID        uint              `gorm:"primaryKey;autoIncrement" json:"id"`
	Email     string            `gorm:"unique;not null" json:"email"`
	FullName  string            `gorm:"size:255;not null" json:"fullName"`
	Password  []byte            `gorm:"not null" json:"-" validate:"required,min=8"`
	Projects  []UserProjectRole `gorm:"foreignKey:UserID" json:"projects"`
	PlanID    uint              `gorm:"not null" json:"planId"`
	Plan      Plan              `gorm:"foreignKey:PlanID" json:"-"`
	CreatedAt time.Time         `json:"createdAt"`
	UpdatedAt time.Time         `json:"updatedAt"`
}
type UserProjectRole struct {
	ID        uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint    `json:"userId"`
	ProjectID uint    `json:"projectId"`
	RoleID    uint    `json:"roleId"`
	User      User    `gorm:"foreignKey:UserID" json:"-"`
	Project   Project `gorm:"foreignKey:ProjectID" json:"-"`
	Role      Role    `gorm:"foreignKey:RoleID" json:"-"`
}

func (user *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return nil
}
func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
