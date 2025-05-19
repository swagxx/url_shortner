package dto

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"'`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
