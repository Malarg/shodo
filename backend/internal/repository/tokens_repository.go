package repository

import (
	"shodo/internal/models"

	"github.com/go-redis/redis"
)

type TokensRepository struct {
	Redis *redis.Client
}

func (repository *TokensRepository) SaveTokens(userId string, tokens *models.AuthTokens) error {
	status := repository.Redis.Set(userId+"_access", tokens.Access, 0)
	if status.Err() != nil {
		return status.Err()
	}

	status = repository.Redis.Set(userId+"_refresh", tokens.Refresh, 0)
	if status.Err() != nil {
		return status.Err()
	}

	return nil
}

func (repository *TokensRepository) GetTokens(userId string) (*models.AuthTokens, error) {
	var tokens models.AuthTokens
	err := repository.Redis.Get(userId).Scan(&tokens)
	if err != nil {
		return nil, err
	}

	return &tokens, nil
}
