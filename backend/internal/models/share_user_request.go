package models

type ShareUserRequest struct {
	UserId string `json:"user_id" binding:"required"`
	ListId string `json:"list_id" binding:"required"`
}
