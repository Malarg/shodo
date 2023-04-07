package repository

import "shodo/models"

type TaskList interface {
	CreateTaskList(taskList *models.TaskList) error
	DeleteTaskList(id string) error
	AddUserToList(listId string, userId string) error
	RemoveUserFromList(listId string, userId string) error
	AddTaskToList(listId *string, task *models.Task) error
	RemoveTaskFromList(listId *string, taskId *string) error
	GetTaskList(id *string) (models.TaskList, error)
}

type Users interface {
	CreateUser(user models.User) (string, error)
	CheckUserExists(email string) (bool, error)
	DeleteUser(id string) error
	GetUserByEmail(email string) (models.User, error)
}
