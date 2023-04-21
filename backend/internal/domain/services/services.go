package services

import (
	"context"
	"shodo/models"
)

type Authentication interface {
	LogIn(request models.LoginUserRequest) (*models.AuthTokens, error)
	IsAuthorized(token string) (bool, error)
}

type Registration interface {
	Register(ctx context.Context, request models.RegisterUserRequest) (*models.AuthTokens, error)
}

type TaskList interface {
	CreateDefaultTaskList(ctx context.Context, username string, ownerId string)
	GetTaskList(listId *string, userToken string) (models.TaskList, *models.Error)
	GetTaskLists(userToken string) ([]models.TaskListShort, error)
	CreateTaskList(ctx context.Context, list *models.TaskList)
	AddTaskToList(listId *string, task *models.Task, userToken string) (*string, *models.Error)
	RemoveTaskFromList(listId *string, taskId *string, userToken string) *models.Error
	IsEditListAllowed(listId *string, userToken string) (bool, error)
	StartShareWithUser(listId *string, email *string, userToken string) *models.Error
	StopShareWithUser(listId *string, email *string, userToken string) *models.Error
}

type Users interface {
	GetAllUsers(userToken string) ([]models.UserShort, error)
}

type Tokens interface {
	GenerateAndSaveTokens(userId *string) (*models.AuthTokens, error)
	GetTokens(userId string) (*models.AuthTokens, error)
}
