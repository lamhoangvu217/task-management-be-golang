package services

import (
	"github.com/lamhoangvu217/task-management-be-golang/database"
	"github.com/lamhoangvu217/task-management-be-golang/models"
)

func GetAllRoles() ([]models.Role, error) {
	var roles []models.Role
	if err := database.DB.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func CreateRole(role *models.Role) (*models.Role, error) {
	if err := database.DB.Create(&role).Error; err != nil {
		return nil, err
	}
	return role, nil
}

func DeleteRole(role *models.Role, roleId uint) error {
	if err := database.DB.Delete(&role, roleId).Error; err != nil {
		return err
	}
	return nil
}

//func AssignRoleToUser() error {
//
//}
