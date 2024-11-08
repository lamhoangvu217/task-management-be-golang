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

func GetProjectByUserId(userId uint) ([]models.Project, error) {
	var projects []models.Project
	query := database.DB.Where("owner_id = ?", userId).Preload("Users")
	if err := query.Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

func AddCollaboratorToProject(res *models.UserProjectRole) (*models.UserProjectRole, error) {
	if err := database.DB.Create(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
