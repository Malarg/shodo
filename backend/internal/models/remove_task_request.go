package models

type RemoveTaskRequest struct {
	Task   Task   `json:"task"`
	ListId string `json:"list_id"`
}
