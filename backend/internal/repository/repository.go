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
	CreateUser(user models.User) (string, error)
	CheckUserExists(email string) (bool, error)
	DeleteUser(id string) error
	GetAllUsers(id string) ([]models.UserShort, error)
	GetUserByEmail(email string) (models.User, error)
	GetUserById(id string) (models.User, error)
}
