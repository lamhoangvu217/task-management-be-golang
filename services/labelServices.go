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
