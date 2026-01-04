package entity

// LoginCredentials represents user login credentials
type LoginCredentials struct {
	Username string `json:"username" example:"admin"`    // Username
	Password string `json:"password" example:"password"` // Password
}
