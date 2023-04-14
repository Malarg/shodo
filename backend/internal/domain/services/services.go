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
	CreateDefaultTaskList(username string, ownerId string)
	GetTaskList(listId *string, userToken string) (models.TaskList, error)
	GetTaskLists(userToken string) ([]models.TaskListShort, error)
	CreateTaskList(list *models.TaskList)
	AddTaskToList(listId *string, task *models.Task, userToken string) (*string, error)
	RemoveTaskFromList(listId *string, taskId *string, userToken string) error
	IsEditListAllowed(listId *string, userToken string) (bool, error)
	StartShareWithUser(listId *string, email *string, userToken string) error
	StopShareWithUser(listId *string, email *string, userToken string) error
}

type Users interface {
	GetAllUsers(userToken string) ([]models.UserShort, error)
}

type Tokens interface {
	GenerateAndSaveTokens(userId *string) (*models.AuthTokens, error)
	GetTokens(userId string) (*models.AuthTokens, error)
}
