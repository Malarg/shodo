package services

import (
	"shodo/models"
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
	GetTaskLists(userToken string) ([]models.TaskListShort, error)
	CreateTaskList(list *models.TaskList)
	AddTaskToList(listId *string, task *models.Task, userToken string) error
	RemoveTaskFromList(listId *string, taskId *string, userToken string) error
	IsEditListAllowed(listId *string, userToken string) (bool, error)
	StartShareWithUser(listId *string, teammateId *string, userToken string) error
	StopShareWithUser(listId *string, teammateId *string, userToken string) error
}

type Tokens interface {
	GenerateAndSaveTokens(userId *string) (*models.AuthTokens, error)
	GetTokens(userId string) (*models.AuthTokens, error)
}
