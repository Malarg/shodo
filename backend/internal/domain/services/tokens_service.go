package services

import (
	"shodo/internal/domain/helpers"
	"shodo/internal/models"
	"shodo/internal/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TokensService struct {
	TokensRepository *repository.TokensRepository
}

func (service *TokensService) GenerateAndSaveTokens(userId primitive.ObjectID) (*models.AuthTokens, error) {
	tokens, err := helpers.GenerateTokens()
	if err != nil {
		return nil, err
	}

	err = service.TokensRepository.SaveTokens(userId.Hex(), &tokens)
	if err != nil {
		return nil, err
	}

	return &models.AuthTokens{Access: tokens.Access, Refresh: tokens.Refresh}, nil
}
