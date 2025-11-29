package models

import (
	"time"

	"github.com/google/uuid"
)

// Assignment представляет задание в курсе
type Assignment struct {
	ID          uuid.UUID  `db:"id" json:"id"`
	CourseID    uuid.UUID  `db:"course_id" json:"course_id"`
	Title       string     `db:"title" json:"title"`
	Description string     `db:"description" json:"description"`
	DueAt       *time.Time `db:"due_at" json:"due_at,omitempty"`
	MaxScore    float64    `db:"max_score" json:"max_score"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at" json:"updated_at"`
}
