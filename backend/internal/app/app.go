package app

import (
	"context"
	"flag"
	"fmt"
	"shodo/internal/config"
	"shodo/internal/domain/services"
	"shodo/internal/repository"
	"shodo/internal/transport"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Run(r *gin.Engine) {
	var configName string
	flag.StringVar(&configName, "configName", "docker", "config file name")
	flag.Parse()
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

	//TODO: add DI?
	tokensRepository := repository.TokensRepository{Redis: rdb}
	usersRepository := repository.UsersRepository{Client: client}
	taskListRepository := repository.TaskListRepository{Mongo: client}

	tasksService := services.TaskListService{TaskListRepository: &taskListRepository}
	tokensService := services.TokensService{TokensRepository: &tokensRepository}
	registrationService := services.RegistrationService{Repository: &usersRepository, TaskListService: &tasksService, TokensService: &tokensService}
	authenticationService := services.AuthenticationService{Repository: &usersRepository, TokensService: &tokensService}

	authHandler := transport.AuthHandler{RegistrationService: &registrationService, AuthenticationService: &authenticationService}
	tasksHandler := transport.TaskListHandler{TaskListService: &tasksService, AuthenticationService: &authenticationService}

	v1 := r.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.LogIn)

		tasks := v1.Group("/tasks")
		// tasks.Use(authHandler.AuthMiddleware())
		tasks.POST("/add", tasksHandler.AddTaskToList)
		tasks.POST("/remove", tasksHandler.DeleteTaskFromList)

		share := v1.Group("/share")
		// share.Use(authHandler.AuthMiddleware())
		share.POST("/start", tasksHandler.StartShareWithUser)
		share.POST("/stop", tasksHandler.StopShareWithUser)
	}
	v1.GET("/ping", transport.PingHandler)

	r.Run()
}

func disconnectMongoDB(ctx context.Context, client *mongo.Client) {
	if err := client.Disconnect(ctx); err != nil {
		panic(err)
	}
}
