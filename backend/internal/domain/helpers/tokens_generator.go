package helpers

import (
	"shodo/internal/models"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	AccessTokenLifeTime  = 3600
	RefreshTokenLifeTime = 86400
	Secret               = "secret" //TODO: investigate and move to env
)

func GenerateTokens() (models.AuthTokens, error) {
	accessToken, err := generateAccessToken()
	if err != nil {
		return models.AuthTokens{}, err
	}
	refreshToken, err := generateRefreshToken()
	if err != nil {
		return models.AuthTokens{}, err
	}
	return models.AuthTokens{Access: accessToken, Refresh: refreshToken}, nil
}

func generateAccessToken() (string, error) {
	return generateToken(AccessTokenLifeTime)
}

func generateRefreshToken() (string, error) {
	return generateToken(RefreshTokenLifeTime)
}

func generateToken(secondsLifetime time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Second * secondsLifetime).Unix(),
	})
	tokenString, err := token.SignedString([]byte(Secret))

	return tokenString, err
}
