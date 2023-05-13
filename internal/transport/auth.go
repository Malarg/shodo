package transport

import (
	"context"
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
func (h *AuthHandler) Register(c *gin.Context) {
	var request models.RegisterUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		h.Logger.Error("Error while binding json", zap.Error(err), zap.Any("request", c.Request.Body))
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), kDefaultTimeout)
	defer cancel()

	tokens, err := h.RegistrationService.Register(ctx, request)
	if err != nil {
		h.Logger.Error("Error while registering user", zap.Error(err), zap.Any("request", request))
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
func (h *AuthHandler) LogIn(c *gin.Context) {
	var request models.LoginUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		h.Logger.Error("Error while binding json", zap.Error(err), zap.Any("request", c.Request.Body))
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), kDefaultTimeout)
	defer cancel()

	tokens, err := h.AuthenticationService.LogIn(ctx, request)
	if err != nil {
		h.Logger.Error("Error while logging in user", zap.Error(err), zap.Any("request", request))
		c.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, tokens)
}

func (h *AuthHandler) LogOut(c *gin.Context) {

}
