package services

import (
	"errors"
	"shodo/internal/domain/helpers"
	"shodo/internal/models"
	"shodo/internal/repository"
)

const (
	kDefaultTaskListTitle = "Shoppings"
	kNotAllowed           = "operation permitted only for list owner or shared users"
)

type TaskListService struct {
	TaskListRepository    repository.TaskList
	AuthenticationService *AuthenticationService
}

func (service *TaskListService) CreateDefaultTaskList(ownerId string) {
	service.TaskListRepository.CreateTaskList(&models.TaskList{Title: kDefaultTaskListTitle, Owner: ownerId})
}

func (service *TaskListService) CreateTaskList(list *models.TaskList) {
	service.TaskListRepository.CreateTaskList(list)
}

func (service *TaskListService) AddTaskToList(listId *string, task *models.Task, userToken string) error {
	isEditListAllowed, err := service.IsEditListAllowed(listId, userToken)

	if err != nil {
		return err
	}

	if !isEditListAllowed {
		return errors.New(kNotAllowed)
	}

	err = service.TaskListRepository.AddTaskToList(listId, task)

	return err
}

func (service *TaskListService) RemoveTaskFromList(listId *string, task *models.Task, userToken string) error {
	isEditListAllowed, err := service.IsEditListAllowed(listId, userToken)

	if err != nil {
		return err
	}

	if !isEditListAllowed {
		return errors.New(kNotAllowed)
	}

	err = service.TaskListRepository.RemoveTaskFromList(listId, task)

	return err
}

func (service *TaskListService) IsEditListAllowed(listId *string, userToken string) (bool, error) {
	list, err := service.TaskListRepository.GetTaskList(listId)
	if err != nil {
		return false, err
	}

	userId, err := helpers.GetUserIdFromToken(userToken)
	if err != nil {
		return false, err
	}

	if userId == list.Owner || helpers.Contains(list.SharedWith, userId) {
		return true, nil
	}

	return false, nil
}
