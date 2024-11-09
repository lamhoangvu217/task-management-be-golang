package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lamhoangvu217/task-management-be-golang/constants"
	"github.com/lamhoangvu217/task-management-be-golang/models"
	"strconv"
	"time"
)

const SecretKey = "secret"

func GenerateJwt(issuer uint, email string) (string, error) {
	claims := jwt.MapClaims{
		"iss":   issuer,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
		"email": email,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ParseJwt(cookie string) (models.User, error) {
	var userInfo models.User
	token, err := jwt.ParseWithClaims(cookie, &jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil || !token.Valid {
		return userInfo, err
	}
	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok {
		return userInfo, fmt.Errorf("invalid token")
	}
	if email, ok := (*claims)["email"].(string); ok {
		userInfo.Email = email
	}
	if userId, ok := (*claims)["iss"].(float64); ok { // Check for float64
		userInfo.ID = uint(userId) // Convert float64 to uint directly
	} else if userIdStr, ok := (*claims)["iss"].(string); ok { // Check for string
		userId, err := strconv.ParseUint(userIdStr, 10, 32)
		if err == nil {
			userInfo.ID = uint(userId)
		}
	}
	return userInfo, nil
}

func IsValidTaskStatus(status string) bool {
	return status == constants.TaskStatusTodo || status == constants.TaskStatusDoing || status == constants.TaskStatusDone
}

func IsValidSubtaskStatus(status string) bool {
	return status == constants.SubtaskStatusTodo || status == constants.SubtaskStatusDoing || status == constants.SubtaskStatusDone
}

func IsValidTaskPriority(priority string) bool {
	return priority == constants.TaskPriorityLow || priority == constants.TaskPriorityMedium || priority == constants.TaskPriorityHigh
}

func IsValidProjectStatus(status string) bool {
	return status == constants.ProjectStatusActive || status == constants.ProjectStatusInActive
}
