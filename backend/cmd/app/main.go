package main

import (
	"shodo/internal/app"

	_ "shodo/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Shodo API
// @version 1.0
// @description This is methods declaration for Shodo API
// @host localhost:8080
// @BasePath /api/v1
func main() {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	app.Run(r)
}
