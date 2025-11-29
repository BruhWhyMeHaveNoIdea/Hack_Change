package dto

import "time"

type SubmissionResponse struct {
	ID           string     `json:"id"`
	AssignmentID string     `json:"assignment_id"`
	StudentID    string     `json:"student_id"`
	Content      string     `json:"content"`
	SubmittedAt  time.Time  `json:"submitted_at"`
	Score        *float64   `json:"score,omitempty"`
	GradedAt     *time.Time `json:"graded_at,omitempty"`
	Feedback     *string    `json:"feedback,omitempty"`
}
