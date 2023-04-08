package transport

import (
	"net/http"
	"shodo/internal/domain/services"
	"shodo/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthHandler struct {
	RegistrationService   services.Registration
	AuthenticationService services.Authentication
	Logger                *zap.Logger
}

// Register godoc
// @Summary Register new user
// @Description Register new user
// @Tags auth
// @Accept  json
// @Produce  json
// @Param user body models.RegisterUserRequest true "User data"
// @Success 200 {object} models.AuthTokens
// @Failure 400 {object} models.Error
// @Router /api/v1/auth/register [post]
func (handler *AuthHandler) Register(c *gin.Context) {
	var request models.RegisterUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		handler.Logger.Error("Error while binding json", zap.Error(err), zap.Any("request", c.Request.Body))
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}

	tokens, err := handler.RegistrationService.Register(request)
	if err != nil {
		handler.Logger.Error("Error while registering user", zap.Error(err), zap.Any("request", request))
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, tokens)
}

// LogIn godoc
// @Summary Log in user
// @Description Log in user
// @Tags auth
// @Accept  json
// @Produce  json
// @Param user body models.LoginUserRequest true "User data"
// @Success 200 {object} models.AuthTokens
// @Failure 400 {object} models.Error
// @Router /api/v1/auth/login [post]
func (handler *AuthHandler) LogIn(c *gin.Context) {
	var request models.LoginUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		handler.Logger.Error("Error while binding json", zap.Error(err), zap.Any("request", c.Request.Body))
		c.JSON(http.StatusBadGateway, models.Error{Message: err.Error()})
		return
	}

	tokens, err := handler.AuthenticationService.LogIn(request)
	if err != nil {
		handler.Logger.Error("Error while logging in user", zap.Error(err), zap.Any("request", request))
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, tokens)
}

func (handler *AuthHandler) LogOut(c *gin.Context) {

}
