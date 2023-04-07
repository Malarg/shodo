package services

import (
	"errors"
	"shodo/internal/domain/helpers"
	"shodo/internal/repository"
	"shodo/models"
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

func (service *TaskListService) RemoveTaskFromList(listId *string, taskId *string, userToken string) error {
	isEditListAllowed, err := service.IsEditListAllowed(listId, userToken)

	if err != nil {
		return err
	}

	if !isEditListAllowed {
		return errors.New(kNotAllowed)
	}

	err = service.TaskListRepository.RemoveTaskFromList(listId, taskId)

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

func (service *TaskListService) StartShareWithUser(listId *string, teammateId *string, userToken string) error {
	userId, err := helpers.GetUserIdFromToken(userToken)
	if err != nil {
		return err
	}

	if userId == *teammateId {
		return errors.New("can't share with yourself")
	}

	list, err := service.TaskListRepository.GetTaskList(listId)
	if err != nil {
		return err
	}

	if list.Owner != userId {
		return errors.New("only list owner can share list")
	}

	if helpers.Contains(list.SharedWith, *teammateId) {
		return errors.New("user already shared")
	}

	err = service.TaskListRepository.AddUserToList(*listId, *teammateId)

	return err
}

func (service *TaskListService) StopShareWithUser(listId *string, teammateId *string, userToken string) error {
	userId, err := helpers.GetUserIdFromToken(userToken)
	if err != nil {
		return err
	}

	if userId == *teammateId {
		return errors.New("can't stop share with yourself")
	}

	list, err := service.TaskListRepository.GetTaskList(listId)
	if err != nil {
		return err
	}

	if list.Owner != userId {
		return errors.New("only list owner can stop share list")
	}

	if !helpers.Contains(list.SharedWith, *teammateId) {
		return errors.New("user not shared")
	}

	err = service.TaskListRepository.RemoveUserFromList(*listId, *teammateId)

	return err
}
