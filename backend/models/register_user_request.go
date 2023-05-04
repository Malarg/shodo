package models

type RegisterUserRequest struct {
	Username string `json:"username" binding:"required,min=3,max=32"`
	Password string `json:"password" binding:"required,min=8,max=32"`
	Email    string `json:"email" binding:"required,email"`
}
