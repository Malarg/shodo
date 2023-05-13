package models

type TaskList struct {
	ID         string   `json:"id"`
	Title      string   `json:"title"`
	Owner      string   `json:"owner"`
	SharedWith []string `json:"shared_with"`
	Tasks      []Task   `json:"tasks"`
}

type TaskListShort struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Owner string `json:"owner"`
}
