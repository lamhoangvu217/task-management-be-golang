package services

import (
	"errors"
	"github.com/lamhoangvu217/task-management-be-golang/database"
	"github.com/lamhoangvu217/task-management-be-golang/models"
	"gorm.io/gorm"
)

func GetTasksByUserId(userId uint) ([]models.Task, error) {
	var tasks []models.Task
	if err := database.DB.Where("user_id = ?", userId).Preload("User").Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func CreateTask(task *models.Task) (*models.Task, error) {
	if err := database.DB.Create(&task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func UpdateTask(updatedTask *models.Task) error {
	if err := database.DB.Save(&updatedTask).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTask(taskId uint) error {
	// Start a transaction
	tx := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Retrieve the task to ensure it exists
	var task models.Task
	if err := tx.First(&task, taskId).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("task id not found")
		}
		return errors.New("could not retrieve task")
	}

	// Delete all subtasks associated with the task
	if err := tx.Where("task_id = ?", taskId).Delete(&models.Subtask{}).Error; err != nil {
		tx.Rollback()
		return errors.New("failed to delete subtasks")
	}

	// Clear the association between the label and tasks
	if err := tx.Model(task).Association("Labels").Clear(); err != nil {
		tx.Rollback()
		return errors.New("failed to remove association between the label and tasks")
	}

	// Delete the task itself
	if err := tx.Delete(&task).Error; err != nil {
		tx.Rollback()
		return errors.New("failed to delete task")
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return errors.New("transaction commit failed")
	}

	return nil
}
