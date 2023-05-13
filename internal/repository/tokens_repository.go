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

func (this *TokensRepository) SaveTokens(userId string, tokens *models.AuthTokens) error {
	status := this.Redis.Set(userId+"_access", tokens.Access, t.AccessTokenLifeTime*time.Second)
	if status.Err() != nil {
		return status.Err()
	}

	status = this.Redis.Set(userId+"_refresh", tokens.Refresh, t.RefreshTokenLifeTime*time.Second)
	if status.Err() != nil {
		return status.Err()
	}

	return nil
}

func (this *TokensRepository) GetTokens(userId string) (*models.AuthTokens, error) {
	access, err := this.Redis.Get(userId + "_access").Result()
	if err != nil {
		return nil, err
	}

	refresh, err := this.Redis.Get(userId + "_refresh").Result()
	if err != nil {
		return nil, err
	}

	return &models.AuthTokens{Access: access, Refresh: refresh}, nil
}
