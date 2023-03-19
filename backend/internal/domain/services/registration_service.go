package services

import (
	"errors"
	"fmt"
	"shodo/internal/domain/helpers"
	"shodo/internal/models"
	"shodo/internal/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	MinPasswordLength = 6
)

type RegistrationService struct {
	Repository    *repository.UsersRepository
	TokensService *TokensService
}

func (service *RegistrationService) Register(request models.RegisterUserRequest) (*models.AuthTokens, error) {
	if err := service.checkIfCanRegister(request); err != nil {
		return nil, err
	}

	userId, err := service.putNewUserToDb(request)
	if err != nil {
		return nil, err
	}

	tokens, err := service.TokensService.GenerateAndSaveTokens(userId)
	if err != nil {
		service.Repository.DeleteUser(userId)
		return nil, err
	}

	return tokens, nil
}

func (service *RegistrationService) checkIfCanRegister(request models.RegisterUserRequest) error {
	userExists, err := service.Repository.CheckUserExists(request.Email)
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

func (service *RegistrationService) putNewUserToDb(request models.RegisterUserRequest) (primitive.ObjectID, error) {
	hashedPassword, err := helpers.HashPassword(request.Password)
	if err != nil {
		return primitive.NilObjectID, err
	}

	user := models.User{
		Email:    request.Email,
		Username: request.Username,
		Password: hashedPassword,
	}

	return service.Repository.CreateUser(user)
}
