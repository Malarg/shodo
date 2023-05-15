package services

import (
	"context"
	"net/http"
	"shodo/internal/domain/tokens"
	"shodo/internal/repository"
	"shodo/models"

	"golang.org/x/exp/slices"
)

const (
	defaultTaskListTitle = "Shoppings"
	notAllowedMessage    = "operation permitted only for list owner or shared users"
)

type TaskListService struct {
	TaskListRepository    repository.TaskList
	UsersRepository       repository.Users
	AuthenticationService *AuthenticationService
}

func (s *TaskListService) CreateDefaultTaskList(ctx context.Context, username, ownerId string) {
	s.TaskListRepository.CreateTaskList(ctx, &models.TaskList{Title: username + " " + defaultTaskListTitle, Owner: ownerId})
}

func (s *TaskListService) CreateTaskList(ctx context.Context, list *models.TaskList) {
	s.TaskListRepository.CreateTaskList(ctx, list)
}

func (s *TaskListService) GetTaskLists(ctx context.Context, userToken string) ([]models.TaskListShort, error) {
	userId, err := tokens.GetUserIdFromToken(userToken)
	if err != nil {
		return nil, err
	}

	return s.TaskListRepository.GetTaskLists(ctx, userId)
}

func (s *TaskListService) GetTaskList(ctx context.Context, listId *string, userToken string) (models.TaskList, *models.Error) {
	isListExists, err := s.TaskListRepository.CheckTaskListExists(ctx, *listId)
	if !isListExists {
		return models.TaskList{}, &models.Error{Code: http.StatusNotFound, Message: "list not found"}
	}

	isEditListAllowed, err := s.IsEditListAllowed(ctx, listId, userToken)
	if err != nil {
		return models.TaskList{}, &models.Error{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	if !isEditListAllowed {
		return models.TaskList{}, &models.Error{Code: http.StatusForbidden, Message: notAllowedMessage}
	}

	resp, err := s.TaskListRepository.GetTaskList(ctx, listId)
	if err != nil {
		return models.TaskList{}, &models.Error{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	return resp, nil
}

func (s *TaskListService) AddTaskToList(ctx context.Context, listId *string, task *models.Task, userToken string) (*string, *models.Error) {
	isListExists, err := s.TaskListRepository.CheckTaskListExists(ctx, *listId)
	if !isListExists {
		return nil, &models.Error{Code: http.StatusNotFound, Message: "list not found"}
	}

	isEditListAllowed, err := s.IsEditListAllowed(ctx, listId, userToken)
	if err != nil {
		return nil, &models.Error{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	if !isEditListAllowed {
		return nil, &models.Error{Code: http.StatusForbidden, Message: notAllowedMessage}
	}

	taskId, err := s.TaskListRepository.AddTaskToList(ctx, listId, task)
	if err != nil {
		return nil, &models.Error{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	return taskId, nil
}

func (s *TaskListService) RemoveTaskFromList(ctx context.Context, listId *string, taskId *string, userToken string) *models.Error {
	isListExists, err := s.TaskListRepository.CheckTaskListExists(ctx, *listId)
	if !isListExists {
		return &models.Error{Code: http.StatusNotFound, Message: "list not found"}
	}

	isTaskExists, err := s.TaskListRepository.CheckTaskExists(ctx, *listId, *taskId)
	if !isTaskExists {
		return &models.Error{Code: http.StatusNotFound, Message: "task not found"}
	}

	isEditListAllowed, err := s.IsEditListAllowed(ctx, listId, userToken)
	if err != nil {
		return &models.Error{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	if !isEditListAllowed {
		return &models.Error{Code: http.StatusForbidden, Message: notAllowedMessage}
	}

	err = s.TaskListRepository.RemoveTaskFromList(ctx, listId, taskId)
	if err != nil {
		return &models.Error{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	return nil
}

func (s *TaskListService) IsEditListAllowed(ctx context.Context, listId *string, userToken string) (bool, error) {
	list, err := s.TaskListRepository.GetTaskList(ctx, listId)
	if err != nil {
		return false, err
	}

	userId, err := tokens.GetUserIdFromToken(userToken)
	if err != nil {
		return false, err
	}

	if userId == list.Owner || slices.Contains(list.SharedWith, userId) {
		return true, nil
	}

	return false, nil
}

func (s *TaskListService) StartShareWithUser(ctx context.Context, listId *string, email *string, userToken string) *models.Error {
	selfId, err := tokens.GetUserIdFromToken(userToken)
	if err != nil {
		return &models.Error{Message: err.Error(), Code: http.StatusBadRequest}
	}

	user, err := s.UsersRepository.GetUserByEmail(ctx, *email)
	if selfId == *&user.ID {
		return &models.Error{Message: "can't share with yourself", Code: http.StatusBadRequest}
	}

	list, err := s.TaskListRepository.GetTaskList(ctx, listId)
	if err != nil {
		return &models.Error{Message: err.Error(), Code: http.StatusBadRequest}
	}

	if list.Owner != selfId {
		return &models.Error{Message: "only list owner can share list", Code: http.StatusForbidden}
	}

	if slices.Contains(list.SharedWith, *&user.ID) {
		return &models.Error{Message: "user already shared", Code: http.StatusBadRequest}
	}

	err = s.TaskListRepository.AddUserToList(ctx, *listId, *email)

	if err != nil {
		return &models.Error{Message: err.Error(), Code: http.StatusBadRequest}
	}

	return nil
}

func (s *TaskListService) StopShareWithUser(ctx context.Context, listId *string, email *string, userToken string) *models.Error {
	selfId, err := tokens.GetUserIdFromToken(userToken)
	if err != nil {
		return &models.Error{Message: err.Error(), Code: http.StatusBadRequest}
	}

	user, err := s.UsersRepository.GetUserByEmail(ctx, *email)
	if selfId == *&user.ID {
		return &models.Error{Message: "can't stop share with yourself", Code: http.StatusBadRequest}
	}

	list, err := s.TaskListRepository.GetTaskList(ctx, listId)
	if err != nil {
		return &models.Error{Message: err.Error(), Code: http.StatusBadRequest}
	}

	if list.Owner != selfId {
		return &models.Error{Message: "only list owner can stop share list", Code: http.StatusForbidden}
	}

	if !slices.Contains(list.SharedWith, user.ID) {
		return &models.Error{Message: "user not shared", Code: http.StatusBadRequest}
	}

	err = s.TaskListRepository.RemoveUserFromList(ctx, *listId, *email)

	if err != nil {
		return &models.Error{Message: err.Error(), Code: http.StatusBadRequest}
	}

	return nil
}
