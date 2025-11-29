package models

import (
	"time"

	"github.com/google/uuid"
)

// Material представляет учебный материал (ссылку/файл/ресурс) в курсе
type Material struct {
	ID          uuid.UUID `db:"id" json:"id"`
	CourseID    uuid.UUID `db:"course_id" json:"course_id"`
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	URL         *string   `db:"url" json:"url,omitempty"`
	StoragePath *string   `db:"storage_path" json:"storage_path,omitempty"`
	UploaderID  uuid.UUID `db:"uploader_id" json:"uploader_id"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}
