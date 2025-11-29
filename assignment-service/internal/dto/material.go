package dto

import "time"

// MaterialCreateRequest — payload для создания материала (multipart file handling handled separately)
type MaterialCreateRequest struct {
	Title       string  `json:"title" binding:"required"`
	Description *string `json:"description,omitempty"`
	URL         *string `json:"url,omitempty"`
}

// MaterialResponse — ответ для материала
type MaterialResponse struct {
	ID          string    `json:"id"`
	CourseID    string    `json:"course_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         *string   `json:"url,omitempty"`
	StoragePath *string   `json:"storage_path,omitempty"`
	UploaderID  string    `json:"uploader_id"`
	CreatedAt   time.Time `json:"created_at"`
}
