package transport

import (
	"net/http"
	"shodo/internal/domain/services"
	"shodo/internal/models"

	"github.com/gin-gonic/gin"
)

//Read params from query and send it to service layer.
//Question: where validation should be? Looks like that at service layer

type AuthHandler struct {
	RegistrationService   *services.RegistrationService
	AuthenticationService *services.AuthenticationService
}

func (handler *AuthHandler) Register(c *gin.Context) {
	var request models.RegisterUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokens, err := handler.RegistrationService.Register(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tokens)
}

func (handler *AuthHandler) LogIn(c *gin.Context) {
	var request models.LoginUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	tokens, err := handler.AuthenticationService.LogIn(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tokens)
}

func (handler *AuthHandler) LogOut(c *gin.Context) {

}
