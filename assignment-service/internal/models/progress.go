package models

import "time"

// Progress соответствует таблице `progress`
type Progress struct {
	ID             int64     `db:"progress_id" json:"progress_id"`
	StudentID      int64     `db:"student_id" json:"student_id"`
	TaskID         int64     `db:"task_id" json:"task_id"`
	Status         string    `db:"status" json:"status"`
	CompletionDate time.Time `db:"completion_date" json:"completion_date"`
	Score          *float32  `db:"score" json:"score,omitempty"`
}
