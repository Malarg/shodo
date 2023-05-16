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

	mongoClient, err := createMongoClient(context.Background(), config)
	if err != nil {
		panic(err)
	}
	defer disconnectMongoDB(mongoClient, logger)

	redisAddr := fmt.Sprintf("%s:6379", config.RedisHost)
	fmt.Println(redisAddr)
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",
		DB:       0,
	})
	defer redisClient.Close()

	app.Run(logger, config, mongoClient, redisClient)
}

func createMongoClient(ctx context.Context, config *config.Config) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MongoUrl))
	if err != nil {
		panic(err)
	}
	return mongoClient, nil
}

func disconnectMongoDB(client *mongo.Client, logger *zap.Logger) {
	if err := client.Disconnect(nil); err != nil {
		logger.Error("Failed to disconnect MongoDB", zap.Error(err))
	}
}
