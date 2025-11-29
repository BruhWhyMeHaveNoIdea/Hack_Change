package models

import (
	"time"
)

// Role перечисляет доступные роли пользователя
type Role string

const (
	RoleStudent Role = "student"
	RoleTeacher Role = "teacher"
	RoleAdmin   Role = "admin"
)

// User представляет пользователя системы (студент, преподаватель, админ)
type User struct {
	ID           string    `db:"student_id" json:"id"`
	Email        string    `db:"email" json:"email"`
	PasswordHash string    `db:"password" json:"-"`
	FirstName    string    `db:"first_name" json:"first_name"`
	LastName     string    `db:"last_name" json:"last_name"`
	MiddleName   *string   `db:"middle_name" json:"middle_name,omitempty"`
	Role         Role      `db:"-" json:"role"`
	CreatedAt    time.Time `db:"-" json:"created_at"`
	UpdatedAt    time.Time `db:"-" json:"updated_at"`
}

// IsStudent возвращает true, если пользователь — студент
func (u *User) IsStudent() bool { return u != nil && u.Role == RoleStudent }

// IsTeacher возвращает true, если пользователь — преподаватель
func (u *User) IsTeacher() bool { return u != nil && u.Role == RoleTeacher }

// IsAdmin возвращает true, если пользователь — админ
func (u *User) IsAdmin() bool { return u != nil && u.Role == RoleAdmin }
