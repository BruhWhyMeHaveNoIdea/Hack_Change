package dto

import "time"

type UserResponse struct {
	ID        string      `json:"id"`
	Email     string      `json:"email"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	MiddleName  *string   `json:"middle_name,omitempty"`
	Role      string      `json:"role"`
	CreatedAt time.Time   `json:"created_at"`
}
