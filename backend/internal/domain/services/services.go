package services

import (
	"shodo/internal/models"
)

//go:generate mockgen -source=services.go -destination=mocks/mock.go

type Authentication interface {
	LogIn(request models.LoginUserRequest) (*models.AuthTokens, error)
	IsAuthorized(token string) (bool, error)
}

type Registration interface {
	Register(request models.RegisterUserRequest) (*models.AuthTokens, error)
}

type TaskList interface {
	CreateDefaultTaskList(ownerId string)
	CreateTaskList(list *models.TaskList)
	AddTaskToList(listId *string, task *models.Task, userToken string) error
	RemoveTaskFromList(listId *string, taskId *string, userToken string) error
	IsEditListAllowed(listId *string, userToken string) (bool, error)
}

type Tokens interface {
	GenerateAndSaveTokens(userId *string) (*models.AuthTokens, error)
	GetTokens(userId string) (*models.AuthTokens, error)
}