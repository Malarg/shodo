package services

import (
	"shodo/internal/domain/tokens"
	"shodo/internal/repository"
	"shodo/models"
)

type TokensService struct {
	TokensRepository *repository.TokensRepository
}

func (s *TokensService) GenerateAndSaveTokens(userId string) (*models.AuthTokens, error) {
	tokens, err := tokens.GenerateTokens(userId)
	if err != nil {
		return nil, err
	}

	err = s.TokensRepository.SaveTokens(userId, &tokens)
	if err != nil {
		return nil, err
	}

	return &models.AuthTokens{Access: tokens.Access, Refresh: tokens.Refresh}, nil
}

func (s *TokensService) GetTokens(userId string) (*models.AuthTokens, error) {
	tokens, err := s.TokensRepository.GetTokens(userId)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}
