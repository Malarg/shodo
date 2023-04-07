package models

type RemoveTaskRequest struct {
	TaskId string `json:"task_id" binding:"required"`
	ListId string `json:"list_id" binding:"required"`
}
