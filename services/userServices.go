package services

import (
	"github.com/lamhoangvu217/task-management-be-golang/database"
	"github.com/lamhoangvu217/task-management-be-golang/models"
)

func GetUserDetail(userEmail string) (models.User, error) {
	var user models.User
	query := database.DB.Where("email = ?", userEmail).First(&user).Preload("Projects")

	if err := query.Find(&user).Error; err != nil {
		return user, err // Return empty user and error if there's an error
	}
	return user, nil
}

func UpdateUserDetail(updatedUserDetail *models.User) error {
	if err := database.DB.Save(&updatedUserDetail).Error; err != nil {
		return err
	}
	return nil
}

func GetUsers() ([]models.User, error) {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
