package dto

import "time"

type ProgressRequest struct {
	StudentID int64    `json:"studentId" binding:"required"`
	TaskID    int64    `json:"taskId" binding:"required"`
	Status    string   `json:"status" binding:"required"`
	Score     *float32 `json:"score"`
}

type ProgressResponse struct {
	ProgressID     int64     `json:"progress_id"`
	StudentID      int64     `json:"student_id"`
	TaskID         int64     `json:"task_id"`
	Status         string    `json:"status"`
	CompletionDate time.Time `json:"completion_date"`
	Score          *float32  `json:"score,omitempty"`
}
