package transport

import (
	"shodo/internal/domain/services"
	"shodo/internal/models"

	"github.com/gin-gonic/gin"
)

type TaskListHandler struct {
	TaskListService       services.TaskList
	AuthenticationService services.Authentication
}

func (handler *TaskListHandler) AddTaskToList(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	isAuthorized, err := handler.AuthenticationService.IsAuthorized(token)
	if !isAuthorized || err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	var request models.AddTaskRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = handler.TaskListService.AddTaskToList(&request.ListId, &request.Task, token)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{})
}

func (handler *TaskListHandler) DeleteTaskFromList(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	isAuthorized, err := handler.AuthenticationService.IsAuthorized(token)
	if !isAuthorized || err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	var request models.RemoveTaskRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = handler.TaskListService.RemoveTaskFromList(&request.ListId, &request.Task, token)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{})
}
