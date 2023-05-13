package main

import (
	"context"
	"fmt"
	"os"
	_ "shodo/docs"
	"shodo/internal/app"
	"shodo/internal/config"
	"time"

	"github.com/go-redis/redis"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

// @title Shodo API
// @version 1.0
// @description This is methods declaration for Shodo API
// @host localhost:8080
// @BasePath /api/v1
func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	configName := os.Getenv("CONFIG_NAME")
	config, err := config.Init(configName)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoUrl := fmt.Sprintf("mongodb://admin:password@%s:27017", config.MongoHost)
	fmt.Println(mongoUrl)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUrl))
	if err != nil {
		panic(err)
	}
	defer disconnectMongoDB(ctx, client)

	redisAddr := fmt.Sprintf("%s:6379", config.RedisHost)
	fmt.Println(redisAddr)
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",
		DB:       0,
	})
	defer rdb.Close()

	app.Run(logger, config, client, rdb)
}

func disconnectMongoDB(ctx context.Context, client *mongo.Client) {
	if err := client.Disconnect(ctx); err != nil {
		panic(err)
	}
}
