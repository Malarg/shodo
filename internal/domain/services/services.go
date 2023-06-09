package services

import (
	"context"
	"shodo/models"
)

type Authentication interface {
	LogIn(ctx context.Context, request models.LoginUserRequest) (*models.AuthTokens, error)
	IsAuthorized(token string) (bool, error)
}

type Registration interface {
	Register(ctx context.Context, request models.RegisterUserRequest) (*models.AuthTokens, error)
}

type TaskList interface {
	CreateDefaultTaskList(ctx context.Context, username string, ownerId string)
	GetTaskList(ctx context.Context, listId *string, userToken string) (models.TaskList, error)
	GetTaskLists(ctx context.Context, userToken string) ([]models.TaskListShort, error)
	CreateTaskList(ctx context.Context, list *models.TaskList)
	AddTaskToList(ctx context.Context, listId *string, task *models.Task, userToken string) (*string, error)
	RemoveTaskFromList(ctx context.Context, listId *string, taskId *string, userToken string) error
	IsEditListAllowed(ctx context.Context, listId *string, userToken string) (bool, error)
	StartShareWithUser(ctx context.Context, listId *string, email *string, userToken string) error
	StopShareWithUser(ctx context.Context, listId *string, email *string, userToken string) error
}

type Tokens interface {
	GenerateAndSaveTokens(userId *string) (*models.AuthTokens, error)
	GetTokens(userId string) (*models.AuthTokens, error)
}
