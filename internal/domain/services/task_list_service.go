package services

import (
	"context"
	"net/http"
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
	UsersRepository       repository.Users
	AuthenticationService *AuthenticationService
}

func (this *TaskListService) CreateDefaultTaskList(ctx context.Context, username, ownerId string) {
	this.TaskListRepository.CreateTaskList(ctx, &models.TaskList{Title: username + " " + kDefaultTaskListTitle, Owner: ownerId})
}

func (this *TaskListService) CreateTaskList(ctx context.Context, list *models.TaskList) {
	this.TaskListRepository.CreateTaskList(ctx, list)
}

func (this *TaskListService) GetTaskLists(ctx context.Context, userToken string) ([]models.TaskListShort, error) {
	userId, err := helpers.GetUserIdFromToken(userToken)
	if err != nil {
		return nil, err
	}

	return this.TaskListRepository.GetTaskLists(ctx, userId)
}

func (this *TaskListService) GetTaskList(ctx context.Context, listId *string, userToken string) (models.TaskList, *models.Error) {
	isListExists, err := this.TaskListRepository.CheckTaskListExists(ctx, *listId)
	if !isListExists {
		return models.TaskList{}, &models.Error{Code: http.StatusNotFound, Message: "list not found"}
	}

	isEditListAllowed, err := this.IsEditListAllowed(ctx, listId, userToken)
	if err != nil {
		return models.TaskList{}, &models.Error{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	if !isEditListAllowed {
		return models.TaskList{}, &models.Error{Code: http.StatusForbidden, Message: kNotAllowed}
	}

	resp, err := this.TaskListRepository.GetTaskList(ctx, listId)
	if err != nil {
		return models.TaskList{}, &models.Error{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	return resp, nil
}

func (this *TaskListService) AddTaskToList(ctx context.Context, listId *string, task *models.Task, userToken string) (*string, *models.Error) {
	isListExists, err := this.TaskListRepository.CheckTaskListExists(ctx, *listId)
	if !isListExists {
		return nil, &models.Error{Code: http.StatusNotFound, Message: "list not found"}
	}

	isEditListAllowed, err := this.IsEditListAllowed(ctx, listId, userToken)
	if err != nil {
		return nil, &models.Error{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	if !isEditListAllowed {
		return nil, &models.Error{Code: http.StatusForbidden, Message: kNotAllowed}
	}

	taskId, err := this.TaskListRepository.AddTaskToList(ctx, listId, task)
	if err != nil {
		return nil, &models.Error{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	return taskId, nil
}

func (this *TaskListService) RemoveTaskFromList(ctx context.Context, listId *string, taskId *string, userToken string) *models.Error {
	isListExists, err := this.TaskListRepository.CheckTaskListExists(ctx, *listId)
	if !isListExists {
		return &models.Error{Code: http.StatusNotFound, Message: "list not found"}
	}

	isTaskExists, err := this.TaskListRepository.CheckTaskExists(ctx, *listId, *taskId)
	if !isTaskExists {
		return &models.Error{Code: http.StatusNotFound, Message: "task not found"}
	}

	isEditListAllowed, err := this.IsEditListAllowed(ctx, listId, userToken)
	if err != nil {
		return &models.Error{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	if !isEditListAllowed {
		return &models.Error{Code: http.StatusForbidden, Message: kNotAllowed}
	}

	err = this.TaskListRepository.RemoveTaskFromList(ctx, listId, taskId)
	if err != nil {
		return &models.Error{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	return nil
}

func (this *TaskListService) IsEditListAllowed(ctx context.Context, listId *string, userToken string) (bool, error) {
	list, err := this.TaskListRepository.GetTaskList(ctx, listId)
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

func (this *TaskListService) StartShareWithUser(ctx context.Context, listId *string, email *string, userToken string) *models.Error {
	selfId, err := helpers.GetUserIdFromToken(userToken)
	if err != nil {
		return &models.Error{Message: err.Error(), Code: http.StatusBadRequest}
	}

	user, err := this.UsersRepository.GetUserByEmail(ctx, *email)
	if selfId == *&user.ID {
		return &models.Error{Message: "can't share with yourself", Code: http.StatusBadRequest}
	}

	list, err := this.TaskListRepository.GetTaskList(ctx, listId)
	if err != nil {
		return &models.Error{Message: err.Error(), Code: http.StatusBadRequest}
	}

	if list.Owner != selfId {
		return &models.Error{Message: "only list owner can share list", Code: http.StatusForbidden}
	}

	if helpers.Contains(list.SharedWith, *&user.ID) {
		return &models.Error{Message: "user already shared", Code: http.StatusBadRequest}
	}

	err = this.TaskListRepository.AddUserToList(ctx, *listId, *email)

	if err != nil {
		return &models.Error{Message: err.Error(), Code: http.StatusBadRequest}
	}

	return nil
}

func (this *TaskListService) StopShareWithUser(ctx context.Context, listId *string, email *string, userToken string) *models.Error {
	selfId, err := helpers.GetUserIdFromToken(userToken)
	if err != nil {
		return &models.Error{Message: err.Error(), Code: http.StatusBadRequest}
	}

	user, err := this.UsersRepository.GetUserByEmail(ctx, *email)
	if selfId == *&user.ID {
		return &models.Error{Message: "can't stop share with yourself", Code: http.StatusBadRequest}
	}

	list, err := this.TaskListRepository.GetTaskList(ctx, listId)
	if err != nil {
		return &models.Error{Message: err.Error(), Code: http.StatusBadRequest}
	}

	if list.Owner != selfId {
		return &models.Error{Message: "only list owner can stop share list", Code: http.StatusForbidden}
	}

	if !helpers.Contains(list.SharedWith, user.ID) {
		return &models.Error{Message: "user not shared", Code: http.StatusBadRequest}
	}

	err = this.TaskListRepository.RemoveUserFromList(ctx, *listId, *email)

	if err != nil {
		return &models.Error{Message: err.Error(), Code: http.StatusBadRequest}
	}

	return nil
}
