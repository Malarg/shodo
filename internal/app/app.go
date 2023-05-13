package app

import (
	"shodo/internal/domain/services"
	"shodo/internal/repository"
	"shodo/internal/transport"

	"github.com/go-redis/redis"

	"shodo/internal/config"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Run(logger *zap.Logger, config *config.Config, client *mongo.Client, rdb *redis.Client) {
	//TODO: add DI?
	tokensRepository := repository.TokensRepository{Redis: rdb}
	usersRepository := repository.UsersRepository{Client: client, Config: config}
	taskListRepository := repository.TaskListRepository{Mongo: client, Config: config}

	tasksService := services.TaskListService{TaskListRepository: &taskListRepository, UsersRepository: &usersRepository}
	tokensService := services.TokensService{TokensRepository: &tokensRepository}
	usersService := services.UsersService{UsersRepository: &usersRepository}
	registrationService := services.RegistrationService{Repository: &usersRepository, TaskListService: &tasksService, TokensService: &tokensService}
	authenticationService := services.AuthenticationService{Repository: &usersRepository, TokensService: &tokensService}

	authHandler := transport.AuthHandler{RegistrationService: &registrationService, AuthenticationService: &authenticationService, Logger: logger}
	tasksHandler := transport.TaskListHandler{TaskListService: &tasksService, AuthenticationService: &authenticationService, Logger: logger}
	usersHandler := transport.UsersHandler{UsersService: &usersService, AuthenticationService: &authenticationService, Logger: logger}

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.LogIn)

		lists := v1.Group("/lists")
		// lists.Use(authHandler.AuthMiddleware())
		lists.GET("/", tasksHandler.GetLists)
		lists.GET("/:id", tasksHandler.GetTaskList)

		tasks := v1.Group("/tasks")
		// tasks.Use(authHandler.AuthMiddleware())
		tasks.POST("/add", tasksHandler.AddTaskToList)
		tasks.POST("/remove", tasksHandler.DeleteTaskFromList)

		share := v1.Group("/share")
		// share.Use(authHandler.AuthMiddleware())
		share.POST("/start", tasksHandler.StartShareWithUser)
		share.POST("/stop", tasksHandler.StopShareWithUser)

		users := v1.Group("/users")
		// users.Use(authHandler.AuthMiddleware())
		users.GET("/", usersHandler.GetAllUsers)
	}
	v1.GET("/ping", transport.PingHandler)

	r.Run()
}
