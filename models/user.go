package models

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID        uint              `gorm:"primaryKey;autoIncrement" json:"id"`
	Email     string            `gorm:"unique;not null" json:"email"`
	FullName  string            `gorm:"size:255;not null" json:"fullName"`
	Password  []byte            `gorm:"not null" json:"password" validate:"required,min=8"`
	Projects  []UserProjectRole `gorm:"foreignKey:UserID" json:"-"`
	CreatedAt time.Time         `json:"createdAt"`
	UpdatedAt time.Time         `json:"updatedAt"`
}
type UserProjectRole struct {
	UserID    uint    `gorm:"primaryKey" json:"userId"`
	ProjectID uint    `gorm:"primaryKey" json:"projectId"`
	RoleID    uint    `json:"roleId"`
	User      User    `gorm:"foreignKey:UserID"`
	Project   Project `gorm:"foreignKey:ProjectID"`
	Role      Role    `gorm:"foreignKey:RoleID"`
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
