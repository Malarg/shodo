package transport

import (
	"context"
	"net/http"
	"shodo/internal/domain/services"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UsersHandler struct {
	AuthenticationService *services.AuthenticationService
	UsersService          *services.UsersService
	Logger                *zap.Logger
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Get all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} models.UserShort
// @Failure 400 {object} models.Error
// @Router /api/v1/users [get]
func (this *UsersHandler) GetAllUsers(c *gin.Context) {
	token := c.Request.Header.Get(Authorization)
	isAuthorized, err := this.AuthenticationService.IsAuthorized(token)
	if !isAuthorized || err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), kDefaultTimeout)
	defer cancel()

	users, err := this.UsersService.GetAllUsers(ctx, token)
	if err != nil {
		this.Logger.Error("Error while getting users", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}
