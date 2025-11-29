package dto

// DTOs для аутентификации
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Email           string  `json:"email" binding:"required,email"`
	Password        string  `json:"password" binding:"required,min=8"`
	PasswordConfirm string  `json:"password_confirm" binding:"required,eqfield=Password"`
	FirstName       string  `json:"first_name" binding:"required"`
	LastName        string  `json:"last_name" binding:"required"`
	MiddleName      *string `json:"middle_name,omitempty"`
	Role            string  `json:"role,omitempty"`
}

type LoginResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}
