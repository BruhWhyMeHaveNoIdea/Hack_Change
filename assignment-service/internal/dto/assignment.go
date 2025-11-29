package dto

import "time"

type AssignmentResponse struct {
	ID          string     `json:"id"`
	CourseID    string     `json:"course_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	DueAt       *time.Time `json:"due_at,omitempty"`
	MaxScore    float64    `json:"max_score"`
}
