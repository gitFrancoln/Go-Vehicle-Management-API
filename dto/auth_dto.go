package dto

type LoginRequest struct {
	UserEmail    string `json:"user_email" binding:"required,email"`
	UserPassword string `json:"user_password" binding:"required"`
}

type ForgotPasswordRequest struct {
	Email       string `json:"email" binding:"omitempty,email"`
	NewPassword string `json:"new_password" binding:"omitempty,min=6"`

	// Alias opcionales
	UserEmail    string `json:"user_email"`
	UserPassword string `json:"user_password"`
}
