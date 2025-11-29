package models

// ProgressSummary represents aggregated progress counts for a student in a course
type ProgressSummary struct {
	Total     int64 `db:"total" json:"total"`
	Completed int64 `db:"completed" json:"completed"`
}
