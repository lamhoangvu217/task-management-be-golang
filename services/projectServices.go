package services

import (
	"github.com/lamhoangvu217/task-management-be-golang/database"
	"github.com/lamhoangvu217/task-management-be-golang/models"
)

func CreateProject(project *models.Project) (*models.Project, error) {
	if err := database.DB.Create(&project).Error; err != nil {
		return nil, err
	}
	return project, nil
}
