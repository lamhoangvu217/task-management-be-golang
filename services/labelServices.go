package services

import (
	"github.com/lamhoangvu217/task-management-be-golang/database"
	"github.com/lamhoangvu217/task-management-be-golang/models"
)

func CreateLabel(label *models.Label) (*models.Label, error) {
	if err := database.DB.Create(&label).Error; err != nil {
		return nil, err
	}
	return label, nil
}

func GetAllLabels() ([]models.Label, error) {
	var labels []models.Label
	if err := database.DB.Find(&labels).Error; err != nil {
		return nil, err
	}
	return labels, nil
}

func DeleteLabel(label *models.Label, labelId uint) error {
	// Clear the association between the label and tasks
	if err := database.DB.Model(label).Association("Tasks").Clear(); err != nil {
		return err
	}
	if err := database.DB.Delete(&label, labelId).Error; err != nil {
		return err
	}
	return nil
}

func UpdateLabel(updatedLabel *models.Label) error {
	if err := database.DB.Save(&updatedLabel).Error; err != nil {
		return err
	}
	return nil
}
