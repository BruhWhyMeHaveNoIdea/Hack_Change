package dto

type FeedbackItem struct {
	Description string `json:"description"`
	Priority    int    `json:"priority"`
	Status      string `json:"status"`
	SkillName   string `json:"skill_name"`
}
