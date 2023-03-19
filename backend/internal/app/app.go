package app

import (
	"context"
	"shodo/internal/domain/services"
	"shodo/internal/repository"
	"shodo/internal/transport"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Run() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	defer disconnectMongoDB(ctx, client)

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	defer rdb.Close()

	tokensRespotiory := repository.TokensRepository{Redis: rdb}
	usersRepository := repository.UsersRepository{Client: client}
	authService := services.AuthService{Repository: &usersRepository, TokensRepository: &tokensRespotiory}
	authHandler := transport.AuthHandler{Service: &authService}

	r := gin.Default()
	r.GET("/ping", transport.PingHandler)
	r.POST("auth/register", authHandler.Register)
	r.POST("auth/login", authHandler.LogIn)
	r.POST("auth/logout", authHandler.LogOut)
	r.Run()
}

func disconnectMongoDB(ctx context.Context, client *mongo.Client) {
	if err := client.Disconnect(ctx); err != nil {
		panic(err)
	}
}
