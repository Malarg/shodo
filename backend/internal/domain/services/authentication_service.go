package services

import (
	"errors"
	"fmt"
	"shodo/internal/domain/helpers"
	"shodo/internal/repository"
	"shodo/models"
)

type AuthenticationService struct {
	Repository    repository.Users
	TokensService *TokensService
}

func (this *AuthenticationService) LogIn(request models.LoginUserRequest) (*models.AuthTokens, error) {
	user, err := this.Repository.GetUserByEmail(request.Email)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("user with email %s not found", request.Email))
	}

	if !helpers.CheckPasswordHash(request.Password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	tokens, err := this.TokensService.GenerateAndSaveTokens(user.ID)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}

func (this *AuthenticationService) IsAuthorized(token string) (bool, error) {
	userId, err := helpers.GetUserIdFromToken(token)
	tokens, err := this.TokensService.GetTokens(userId)
	if err != nil {
		return false, err
	}

	return tokens.Access == token, nil
}
