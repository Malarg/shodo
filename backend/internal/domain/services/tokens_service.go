package services

import (
	"shodo/internal/domain/helpers"
	"shodo/internal/repository"
	"shodo/models"
)

type TokensService struct {
	TokensRepository *repository.TokensRepository
}

func (this *TokensService) GenerateAndSaveTokens(userId string) (*models.AuthTokens, error) {
	tokens, err := helpers.GenerateTokens(userId)
	if err != nil {
		return nil, err
	}

	err = this.TokensRepository.SaveTokens(userId, &tokens)
	if err != nil {
		return nil, err
	}

	return &models.AuthTokens{Access: tokens.Access, Refresh: tokens.Refresh}, nil
}

func (this *TokensService) GetTokens(userId string) (*models.AuthTokens, error) {
	tokens, err := this.TokensRepository.GetTokens(userId)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}
