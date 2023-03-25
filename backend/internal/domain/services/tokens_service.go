package services

import (
	"shodo/internal/domain/helpers"
	"shodo/internal/models"
	"shodo/internal/repository"
)

type TokensService struct {
	TokensRepository *repository.TokensRepository
}

func (service *TokensService) GenerateAndSaveTokens(userId string) (*models.AuthTokens, error) {
	tokens, err := helpers.GenerateTokens(userId)
	if err != nil {
		return nil, err
	}

	err = service.TokensRepository.SaveTokens(userId, &tokens)
	if err != nil {
		return nil, err
	}

	return &models.AuthTokens{Access: tokens.Access, Refresh: tokens.Refresh}, nil
}

func (service *TokensService) GetTokens(userId string) (*models.AuthTokens, error) {
	tokens, err := service.TokensRepository.GetTokens(userId)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}
