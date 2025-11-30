package dto

import "time"

// AssignmentResponse соответствует загруженному файлу (таблица assignments)
type AssignmentResponse struct {
	AssignmentID string    `json:"assignment_id"`
	StudentID    string    `json:"student_id"`
	TaskID       string    `json:"task_id"`
	UploadDate   time.Time `json:"upload_date"`
	FilePath     *string   `json:"file_path,omitempty"`
	Status       string    `json:"status"`
}
