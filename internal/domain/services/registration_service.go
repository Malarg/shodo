package services

import (
	"errors"
	"fmt"
	"shodo/internal/repository"
	"shodo/models"

	"golang.org/x/crypto/bcrypt"

	"golang.org/x/net/context"
)

const (
	MinPasswordLength = 6
)

type RegistrationService struct {
	Repository      repository.Users
	TaskListService *TaskListService
	TokensService   *TokensService
}

func (this *RegistrationService) Register(ctx context.Context, request models.RegisterUserRequest) (*models.AuthTokens, error) {
	if err := this.checkIfCanRegister(ctx, request); err != nil {
		return nil, err
	}

	userId, err := this.createUser(ctx, request)
	if err != nil {
		return nil, err
	}

	tokens, err := this.TokensService.GenerateAndSaveTokens(userId)
	if err != nil {
		this.Repository.DeleteUser(ctx, userId)
		return nil, err
	}

	this.TaskListService.CreateDefaultTaskList(ctx, request.Username, userId)
	return tokens, nil
}

func (this *RegistrationService) checkIfCanRegister(ctx context.Context, request models.RegisterUserRequest) error {
	userExists, err := this.Repository.CheckUserExists(ctx, request.Email)
	if err != nil {
		return err
	}

	if userExists {
		return errors.New(fmt.Sprintf("user with email %s already exists", request.Email))
	}

	if len(request.Password) < MinPasswordLength {
		message := fmt.Sprintf("password length must be greater than %d characters", MinPasswordLength)
		return errors.New(message)
	}
	return nil
}

func (this *RegistrationService) createUser(ctx context.Context, request models.RegisterUserRequest) (string, error) {
	hashedPassword, err := hashPassword(request.Password)
	if err != nil {
		return "", err
	}

	user := models.User{
		Email:    request.Email,
		Username: request.Username,
		Password: hashedPassword,
	}

	return this.Repository.CreateUser(ctx, user)
}

const (
	DefaultCost int = 10
)

func hashPassword(password string) (string, error) {
	if len(password) > 36 {
		ErrPasswordTooLong := errors.New("password length should not exceed 72 bytes")
		return "", ErrPasswordTooLong
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), DefaultCost)
	return string(bytes), err
}
