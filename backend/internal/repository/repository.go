package repository

import "shodo/models"

type TaskList interface {
	CreateTaskList(taskList *models.TaskList) error
	DeleteTaskList(id string) error
	AddUserToList(listId string, email string) error
	RemoveUserFromList(listId string, email string) error
	AddTaskToList(listId *string, task *models.Task) (*string, error)
	RemoveTaskFromList(listId *string, taskId *string) error
	GetTaskList(id *string) (models.TaskList, error)
	GetTaskLists(userId string) ([]models.TaskListShort, error)
	CheckTaskListExists(id string) (bool, error)
	CheckTaskExists(listId string, taskId string) (bool, error)
}

type Users interface {
	CreateUser(user models.User) (string, error)
	CheckUserExists(email string) (bool, error)
	DeleteUser(id string) error
	GetAllUsers(id string) ([]models.UserShort, error)
	GetUserByEmail(email string) (models.User, error)
	GetUserById(id string) (models.User, error)
}
