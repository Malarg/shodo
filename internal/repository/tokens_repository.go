package repository

import (
	t "shodo/internal/domain/tokens"
	"shodo/models"
	"time"

	"github.com/go-redis/redis"
)

type TokensRepository struct {
	Redis *redis.Client
}

func (r *TokensRepository) SaveTokens(userId string, tokens *models.AuthTokens) error {
	status := r.Redis.Set(userId+"_access", tokens.Access, t.AccessTokenLifeTime*time.Second)
	if status.Err() != nil {
		return status.Err()
	}

	status = r.Redis.Set(userId+"_refresh", tokens.Refresh, t.RefreshTokenLifeTime*time.Second)
	if status.Err() != nil {
		return status.Err()
	}

	return nil
}

func (r *TokensRepository) GetTokens(userId string) (*models.AuthTokens, error) {
	access, err := r.Redis.Get(userId + "_access").Result()
	if err != nil {
		return nil, err
	}

	refresh, err := r.Redis.Get(userId + "_refresh").Result()
	if err != nil {
		return nil, err
	}

	return &models.AuthTokens{Access: access, Refresh: refresh}, nil
}
