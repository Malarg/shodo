package transport

import (
	"shodo/internal/domain/services"
	"shodo/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TaskListHandler struct {
	TaskListService       services.TaskList
	AuthenticationService services.Authentication
	Logger                *zap.Logger
}

// GetLists godoc
// @Summary Get all lists for a user
// @Description Get all lists for a user
// @Tags lists
// @Accept  json
// @Produce  json
// @Success 200 {object} models.GetListsResponse
// @Failure 400 {object} models.Error
// @Router /api/v1/lists [get]
func (handler *TaskListHandler) GetLists(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	isAuthorized, err := handler.AuthenticationService.IsAuthorized(token)
	if !isAuthorized || err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	lists, err := handler.TaskListService.GetTaskLists(token)
	if err != nil {
		handler.Logger.Error("Error while getting lists", zap.Error(err))
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, models.GetListsResponse{Lists: lists})
}

// AddTaskToList godoc
// @Summary Add task to list
// @Description Add task to list
// @Tags lists
// @Accept  json
// @Produce  json
// @Param data body models.AddTaskRequest true "Task data"
// @Success 200 {object} models.IdResponse
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
		handler.Logger.Error("Error while binding add task json", zap.Error(err), zap.Any("request", c.Request.Body))
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := handler.TaskListService.AddTaskToList(&request.ListId, &request.Task, token)
	if err != nil {
		handler.Logger.Error("Error while adding task to list", zap.Error(err), zap.Any("request", request))
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, models.IdResponse{Id: *res})
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
		handler.Logger.Error("Error while binding delete task json", zap.Error(err), zap.Any("request", c.Request.Body))
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = handler.TaskListService.RemoveTaskFromList(&request.ListId, &request.TaskId, token)
	if err != nil {
		handler.Logger.Error("Error while removing task from list", zap.Error(err), zap.Any("request", request))
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{})
}

// GetTaskList godoc
// @Summary Get tasks by list id
// @Description Get tasks by list id
// @Tags lists
// @Accept  json
// @Produce  json
// @Success 200 {object} models.GetTaskListResponse
// @Failure 400 {object} models.Error
// @Router /api/v1/list/:id [post]
func (handler *TaskListHandler) GetTaskList(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	isAuthorized, err := handler.AuthenticationService.IsAuthorized(token)
	if !isAuthorized || err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	listId := c.Param("id")

	list, err := handler.TaskListService.GetTaskList(&listId, token)
	if err != nil {
		handler.Logger.Error("Error while getting list", zap.Error(err), zap.Any("listId", listId))
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, models.GetTaskListResponse{List: list})
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
		handler.Logger.Error("Error while binding start share json", zap.Error(err), zap.Any("request", c.Request.Body))
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = handler.TaskListService.StartShareWithUser(&request.ListId, &request.UserId, token)
	if err != nil {
		handler.Logger.Error("Error while sharing list with user", zap.Error(err), zap.Any("request", request))
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
		handler.Logger.Error("Error while binding sop share json", zap.Error(err), zap.Any("request", c.Request.Body))
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = handler.TaskListService.StopShareWithUser(&request.ListId, &request.UserId, token)
	if err != nil {
		handler.Logger.Error("Error while stop share list with user", zap.Error(err), zap.Any("request", request))
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{})
}
