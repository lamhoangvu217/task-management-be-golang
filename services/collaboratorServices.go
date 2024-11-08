package services

import (
	"github.com/lamhoangvu217/task-management-be-golang/database"
	"github.com/lamhoangvu217/task-management-be-golang/models"
)

func GetCollaboratorsByProjectId(projectId uint) ([]models.CollaboratorResponse, error) {
	var userProjectRoles []models.UserProjectRole
	var collaborators []models.CollaboratorResponse

	// Query the database with Preload for related data
	err := database.DB.Where("project_id = ?", projectId).
		Preload("User").
		Preload("Role").
		Find(&userProjectRoles).Error
	if err != nil {
		return nil, err
	}

	// Map to CollaboratorResponse format
	for _, upr := range userProjectRoles {
		collaborator := models.CollaboratorResponse{
			UserID:   upr.User.ID,
			Email:    upr.User.Email,
			FullName: upr.User.FullName,
			Role:     upr.Role.Name,
		}
		collaborators = append(collaborators, collaborator)
	}

	return collaborators, nil
}

func UpdateCollaboratorInProject(projectId uint, userId uint) error {
	if err := database.DB.Where("project_id = ? AND user_id = ?", projectId, userId).Delete(&models.UserProjectRole{}).Error; err != nil {
		return err
	}
	return nil
}
