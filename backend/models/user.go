package models

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserShort struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}
