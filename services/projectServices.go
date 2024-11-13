package services

import (
	"fmt"
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

func GetUserProjectById(projectId uint, userId uint) (*models.Project, error) {
	var project models.Project
	if err := database.DB.First(&project, projectId).Error; err != nil {
		return nil, err
	}
	if project.OwnerID != userId {
		return nil, fmt.Errorf("you are not owner of this project")
	}
	return &project, nil
}
