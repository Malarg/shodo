package models

type AuthResponse struct {
	AccessToken  string
	RefreshToken string
	ExpiresIn    int
}
