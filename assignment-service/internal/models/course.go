package models

import (
	"time"

	"github.com/google/uuid"
)

// Course представляет учебный курс
type Course struct {
	ID          uuid.UUID `db:"id" json:"id"`
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	TeacherID   uuid.UUID `db:"teacher_id" json:"teacher_id"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
