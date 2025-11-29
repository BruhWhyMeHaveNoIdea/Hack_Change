package models

// FeedbackWithSkill is a projection combining feedback_tickets and competency skill name
type FeedbackWithSkill struct {
	Description string `db:"description" json:"description"`
	Priority    int    `db:"priority" json:"priority"`
	Status      string `db:"status" json:"status"`
	SkillName   string `db:"skill_name" json:"skill_name"`
}
