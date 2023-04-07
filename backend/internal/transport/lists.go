package transport

import (
	"shodo/internal/domain/services"
	"shodo/models"

	"github.com/gin-gonic/gin"
)

type TaskListHandler struct {
	TaskListService       services.TaskList
	AuthenticationService services.Authentication
}

// AddTaskToList godoc
// @Summary Add task to list
// @Description Add task to list
// @Tags lists
// @Accept  json
// @Produce  json
// @Param data body models.AddTaskRequest true "Task data"
// @Success 200 {object} models.EmptyResponse
// @Failure 400 {object} models.Error
// @Router /api/v1/tasks/add [post]
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

// DeleteTaskFromList godoc
// @Summary Delete task from list
// @Description Delete task from list
// @Tags lists
// @Accept  json
// @Produce  json
// @Param data body models.RemoveTaskRequest true "Task data"
// @Success 200 {object} models.EmptyResponse
// @Failure 400 {object} models.Error
// @Router /api/v1/tasks/remove [post]
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

	err = handler.TaskListService.RemoveTaskFromList(&request.ListId, &request.TaskId, token)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{})
}

// StartShareWithUser godoc
// @Summary Start share list with user
// @Description Start share list with user
// @Tags lists
// @Accept  json
// @Produce  json
// @Param data body models.ShareUserRequest true "Share data"
// @Success 200 {object} models.EmptyResponse
// @Failure 400 {object} models.Error
// @Router /api/v1/lists/share/start [post]
func (handler *TaskListHandler) StartShareWithUser(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	isAuthorized, err := handler.AuthenticationService.IsAuthorized(token)
	if !isAuthorized || err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	var request models.ShareUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = handler.TaskListService.StartShareWithUser(&request.ListId, &request.UserId, token)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{})
}

// StopShareWithUser godoc
// @Summary Stop share list with user
// @Description Stop share list with user
// @Tags lists
// @Accept  json
// @Produce  json
// @Param data body models.ShareUserRequest true "Share data"
// @Success 200 {object} models.EmptyResponse
// @Failure 400 {object} models.Error
// @Router /api/v1/lists/share/stop [post]
func (handler *TaskListHandler) StopShareWithUser(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	isAuthorized, err := handler.AuthenticationService.IsAuthorized(token)
	if !isAuthorized || err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	var request models.ShareUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = handler.TaskListService.StopShareWithUser(&request.ListId, &request.UserId, token)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{})
}
