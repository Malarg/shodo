package services

import (
	"errors"
	"fmt"
	"shodo/internal/domain/tokens"
	"shodo/internal/repository"
	"shodo/models"

	"golang.org/x/crypto/bcrypt"

	"golang.org/x/net/context"
)

type AuthenticationService struct {
	Repository    repository.Users
	TokensService *TokensService
}

func (this *AuthenticationService) LogIn(ctx context.Context, request models.LoginUserRequest) (*models.AuthTokens, error) {
	user, err := this.Repository.GetUserByEmail(ctx, request.Email)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("user with email %s not found", request.Email))
	}

	if !checkPasswordHash(request.Password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	tokens, err := this.TokensService.GenerateAndSaveTokens(user.ID)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (this *AuthenticationService) IsAuthorized(token string) (bool, error) {
	userId, err := tokens.GetUserIdFromToken(token)
	tokens, err := this.TokensService.GetTokens(userId)
	if err != nil {
		return false, err
	}

	return tokens.Access == token, nil
}
