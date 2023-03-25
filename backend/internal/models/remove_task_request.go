package models

type RemoveTaskRequest struct {
	TaskId string `json:"task_id"`
	ListId string `json:"list_id"`
}
