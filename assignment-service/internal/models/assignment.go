package models

import (
	"time"
)

// Assignment соответствует таблице `assignments` из миграций — это загруженные работы
type Assignment struct {
	ID         int64     `db:"assignment_id" json:"assignment_id"`
	StudentID  string    `db:"student_id" json:"student_id"`
	TaskID     string    `db:"task_id" json:"task_id"`
	UploadDate time.Time `db:"upload_date" json:"upload_date"`
	FilePath   *string   `db:"file_path" json:"file_path,omitempty"`
	Status     string    `db:"status" json:"status"`
}
