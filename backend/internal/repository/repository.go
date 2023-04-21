package repository

import (
	"context"
	"shodo/models"
)

type TaskList interface {
	CreateTaskList(ctx context.Context, taskList *models.TaskList) error
	DeleteTaskList(ctx context.Context, id string) error
	AddUserToList(ctx context.Context, listId string, email string) error
	RemoveUserFromList(ctx context.Context, listId string, email string) error
	AddTaskToList(ctx context.Context, listId *string, task *models.Task) (*string, error)
	RemoveTaskFromList(ctx context.Context, listId *string, taskId *string) error
	GetTaskList(ctx context.Context, id *string) (models.TaskList, error)
	GetTaskLists(ctx context.Context, userId string) ([]models.TaskListShort, error)
	CheckTaskListExists(ctx context.Context, id string) (bool, error)
	CheckTaskExists(ctx context.Context, listId string, taskId string) (bool, error)
}

type Users interface {
	CreateUser(ctx context.Context, user models.User) (string, error)
	CheckUserExists(ctx context.Context, email string) (bool, error)
	DeleteUser(ctx context.Context, id string) error
	GetAllUsers(ctx context.Context, id string) ([]models.UserShort, error)
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	GetUserById(ctx context.Context, id string) (models.User, error)
}
