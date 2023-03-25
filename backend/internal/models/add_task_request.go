package models

type AddTaskRequest struct {
	Task   Task   `json:"task"`
	ListId string `json:"list_id"`
}
