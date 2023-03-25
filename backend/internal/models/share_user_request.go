package models

type ShareUserRequest struct {
	UserId string `json:"user_id"`
	ListId string `json:"list_id"`
}
