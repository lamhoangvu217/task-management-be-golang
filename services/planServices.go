package services

import (
	"github.com/lamhoangvu217/task-management-be-golang/database"
	"github.com/lamhoangvu217/task-management-be-golang/models"
)

func GetAllPlans() ([]models.Plan, error) {
	var plans []models.Plan
	if err := database.DB.Find(&plans).Error; err != nil {
		return nil, err
	}
	return plans, nil
}

func CreatePlan(plan *models.Plan) (*models.Plan, error) {
	if err := database.DB.Create(&plan).Error; err != nil {
		return nil, err
	}
	return plan, nil
}

func DeletePlan(plan *models.Plan, planId uint) error {
	if err := database.DB.Delete(&plan, planId).Error; err != nil {
		return err
	}
	return nil
}

func UpdatePlan(updatedPlan *models.Plan) error {
	if err := database.DB.Save(&updatedPlan).Error; err != nil {
		return err
	}
	return nil
}

func SubscribePlan(updatedUserPlan *models.User) error {
	if err := database.DB.Save(&updatedUserPlan).Error; err != nil {
		return err
	}
	return nil
}

func GetCurrentUserPlan(user *models.User, userId uint) error {
	if err := database.DB.Preload("Plan").First(&user, userId).Error; err != nil {
		return err
	}
	return nil
}
