package dto

// ProgressSummaryResponse is returned by GET /progress/summary
type ProgressSummaryResponse struct {
	Total     int64 `json:"total"`
	Completed int64 `json:"completed"`
}
