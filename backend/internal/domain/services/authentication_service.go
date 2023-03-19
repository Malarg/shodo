package services

import (
	"errors"
	"fmt"
	"shodo/internal/domain/helpers"
	"shodo/internal/models"
	"shodo/internal/repository"
)

type AuthenticationService struct {
	Repository    *repository.UsersRepository
	TokensService *TokensService
}

func (service *AuthenticationService) LogIn(request models.LoginUserRequest) (*models.AuthTokens, error) {
	user, err := service.Repository.GetUserByEmail(request.Email)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("user with email %s not found", request.Email))
	}

	if !helpers.CheckPasswordHash(request.Password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	tokens, err := service.TokensService.GenerateAndSaveTokens(user.ID)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}
