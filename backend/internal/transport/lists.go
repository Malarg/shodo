package transport

import (
	"net/http"
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

const (
	Authorization = "Authorization"
)

// GetLists godoc
// @Summary Get all lists for a user
// @Description Get all lists for a user
// @Tags lists
// @Accept  json
// @Produce  json
// @Success 200 {object} models.GetListsResponse
// @Failure 400 {object} models.Error
// @Router /api/v1/lists [get]
func (this *TaskListHandler) GetLists(c *gin.Context) {
	token := c.Request.Header.Get(Authorization)
	isAuthorized, err := this.AuthenticationService.IsAuthorized(token)
	if !isAuthorized || err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	lists, err := this.TaskListService.GetTaskLists(token)
	if err != nil {
		this.Logger.Error("Error while getting lists", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.GetListsResponse{Lists: lists})
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
func (this *TaskListHandler) AddTaskToList(c *gin.Context) {
	token := c.Request.Header.Get(Authorization)
	isAuthorized, err := this.AuthenticationService.IsAuthorized(token)
	if !isAuthorized || err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var request models.AddTaskRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		this.Logger.Error("Error while binding add task json", zap.Error(err), zap.Any("request", c.Request.Body))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := this.TaskListService.AddTaskToList(&request.ListId, &request.Task, token)
	if err != nil {
		this.Logger.Error("Error while adding task to list", zap.Error(err), zap.Any("request", request))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.IdResponse{Id: *res})
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
func (this *TaskListHandler) DeleteTaskFromList(c *gin.Context) {
	token := c.Request.Header.Get(Authorization)
	isAuthorized, err := this.AuthenticationService.IsAuthorized(token)
	if !isAuthorized || err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var request models.RemoveTaskRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		this.Logger.Error("Error while binding delete task json", zap.Error(err), zap.Any("request", c.Request.Body))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = this.TaskListService.RemoveTaskFromList(&request.ListId, &request.TaskId, token)
	if err != nil {
		this.Logger.Error("Error while removing task from list", zap.Error(err), zap.Any("request", request))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
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
func (this *TaskListHandler) GetTaskList(c *gin.Context) {
	token := c.Request.Header.Get(Authorization)
	isAuthorized, err := this.AuthenticationService.IsAuthorized(token)
	if !isAuthorized || err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	listId := c.Param("id")

	list, err := this.TaskListService.GetTaskList(&listId, token)
	if err != nil {
		this.Logger.Error("Error while getting list", zap.Error(err), zap.Any("listId", listId))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.GetTaskListResponse{List: list})
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
func (this *TaskListHandler) StartShareWithUser(c *gin.Context) {
	token := c.Request.Header.Get(Authorization)
	isAuthorized, err := this.AuthenticationService.IsAuthorized(token)
	if !isAuthorized || err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var request models.ShareUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		this.Logger.Error("Error while binding start share json", zap.Error(err), zap.Any("request", c.Request.Body))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = this.TaskListService.StartShareWithUser(&request.ListId, &request.UserId, token)
	if err != nil {
		this.Logger.Error("Error while sharing list with user", zap.Error(err), zap.Any("request", request))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
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
func (this *TaskListHandler) StopShareWithUser(c *gin.Context) {
	token := c.Request.Header.Get(Authorization)
	isAuthorized, err := this.AuthenticationService.IsAuthorized(token)
	if !isAuthorized || err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var request models.ShareUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		this.Logger.Error("Error while binding sop share json", zap.Error(err), zap.Any("request", c.Request.Body))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = this.TaskListService.StopShareWithUser(&request.ListId, &request.UserId, token)
	if err != nil {
		this.Logger.Error("Error while stop share list with user", zap.Error(err), zap.Any("request", request))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
