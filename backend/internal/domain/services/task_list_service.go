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

func (this *TaskListService) CreateDefaultTaskList(username, ownerId string) {
	this.TaskListRepository.CreateTaskList(&models.TaskList{Title: username + " " + kDefaultTaskListTitle, Owner: ownerId})
}

func (this *TaskListService) CreateTaskList(list *models.TaskList) {
	this.TaskListRepository.CreateTaskList(list)
}

func (this *TaskListService) GetTaskLists(userToken string) ([]models.TaskListShort, error) {
	userId, err := helpers.GetUserIdFromToken(userToken)
	if err != nil {
		return nil, err
	}

	return this.TaskListRepository.GetTaskLists(userId)
}

func (this *TaskListService) GetTaskList(listId *string, userToken string) (models.TaskList, error) {
	isEditListAllowed, err := this.IsEditListAllowed(listId, userToken)
	if err != nil {
		return models.TaskList{}, err
	}

	if !isEditListAllowed {
		return models.TaskList{}, errors.New(kNotAllowed)
	}

	return this.TaskListRepository.GetTaskList(listId)
}

func (this *TaskListService) AddTaskToList(listId *string, task *models.Task, userToken string) (*string, error) {
	isEditListAllowed, err := this.IsEditListAllowed(listId, userToken)

	if err != nil {
		return nil, err
	}

	if !isEditListAllowed {
		return nil, errors.New(kNotAllowed)
	}

	return this.TaskListRepository.AddTaskToList(listId, task)
}

func (this *TaskListService) RemoveTaskFromList(listId *string, taskId *string, userToken string) error {
	isEditListAllowed, err := this.IsEditListAllowed(listId, userToken)

	if err != nil {
		return err
	}

	if !isEditListAllowed {
		return errors.New(kNotAllowed)
	}

	err = this.TaskListRepository.RemoveTaskFromList(listId, taskId)

	return err
}

func (this *TaskListService) IsEditListAllowed(listId *string, userToken string) (bool, error) {
	list, err := this.TaskListRepository.GetTaskList(listId)
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

func (this *TaskListService) StartShareWithUser(listId *string, teammateId *string, userToken string) error {
	userId, err := helpers.GetUserIdFromToken(userToken)
	if err != nil {
		return err
	}

	if userId == *teammateId {
		return errors.New("can't share with yourself")
	}

	list, err := this.TaskListRepository.GetTaskList(listId)
	if err != nil {
		return err
	}

	if list.Owner != userId {
		return errors.New("only list owner can share list")
	}

	if helpers.Contains(list.SharedWith, *teammateId) {
		return errors.New("user already shared")
	}

	err = this.TaskListRepository.AddUserToList(*listId, *teammateId)

	return err
}

func (this *TaskListService) StopShareWithUser(listId *string, teammateId *string, userToken string) error {
	userId, err := helpers.GetUserIdFromToken(userToken)
	if err != nil {
		return err
	}

	if userId == *teammateId {
		return errors.New("can't stop share with yourself")
	}

	list, err := this.TaskListRepository.GetTaskList(listId)
	if err != nil {
		return err
	}

	if list.Owner != userId {
		return errors.New("only list owner can stop share list")
	}

	if !helpers.Contains(list.SharedWith, *teammateId) {
		return errors.New("user not shared")
	}

	err = this.TaskListRepository.RemoveUserFromList(*listId, *teammateId)

	return err
}
