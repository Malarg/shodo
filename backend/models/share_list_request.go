package models

type ShareListRequest struct {
	Email  string `json:"email" binding:"required"`
	ListId string `json:"list_id" binding:"required"`
}
