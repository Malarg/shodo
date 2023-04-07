package helpers

import (
	"shodo/models"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	AccessTokenLifeTime  = 3600
	RefreshTokenLifeTime = 86400
	Secret               = "aWeSoMeSecReTStRing@3!" //TODO: investigate and move to env
)

func GenerateTokens(userId string) (models.AuthTokens, error) {
	accessToken, err := generateAccessToken(userId)
	if err != nil {
		return models.AuthTokens{}, err
	}
	refreshToken, err := generateRefreshToken(userId)
	if err != nil {
		return models.AuthTokens{}, err
	}
	return models.AuthTokens{Access: accessToken, Refresh: refreshToken}, nil
}

func generateAccessToken(userId string) (string, error) {
	return generateToken(AccessTokenLifeTime, userId)
}

func generateRefreshToken(userId string) (string, error) {
	return generateToken(RefreshTokenLifeTime, userId)
}

func generateToken(secondsLifetime time.Duration, userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": userId,
		"exp": time.Now().Add(time.Second * secondsLifetime).Unix(),
	})
	tokenString, err := token.SignedString([]byte(Secret))

	return tokenString, err
}

func GetUserIdFromToken(token string) (string, error) {
	claims, err := parseToken(token)
	if err != nil {
		return "", err
	}
	return claims["uid"].(string), nil
}

func parseToken(token string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	return claims, err
}
