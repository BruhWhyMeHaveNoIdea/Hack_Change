package dto

import "time"

type CourseResponse struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	TeacherID   string    `json:"teacher_id"`
	CreatedAt   time.Time `json:"created_at"`
}

