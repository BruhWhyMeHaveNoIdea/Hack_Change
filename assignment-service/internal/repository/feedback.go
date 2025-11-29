package repository

import (
	"context"
	"database/sql"
	"fmt"
	"hack_change/internal/models"
	postgres "hack_change/pkg/db"
	"hack_change/pkg/logger"

	sq "github.com/Masterminds/squirrel"
)

type FeedbackRepo struct {
	db     *postgres.DB
	logger logger.Logger
}

func NewFeedbackRepository(db *postgres.DB) *FeedbackRepo {
	return &FeedbackRepo{db: db, logger: logger.New("feedback_repo")}
}

// GetByAssignmentID returns feedback items and associated competency skill name for an assignment
func (r *FeedbackRepo) GetByAssignmentID(ctx context.Context, assignmentID int64) ([]*models.FeedbackWithSkill, error) {
	qb := sq.Select("ft.description", "ft.priority", "ft.status", "c.skill_name").
		From("feedback_tickets ft").
		Join("competencies c ON ft.competency_id = c.competency_id").
		Where(sq.Eq{"ft.assignment_id": assignmentID}).
		PlaceholderFormat(sq.Dollar)

	sqlStr, args, err := qb.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build feedback query: %w", err)
	}

	rows, err := r.db.Db.QueryContext(ctx, sqlStr, args...)
	if err != nil {
		r.logger.Error(ctx, "failed to query feedback", "error", err)
		return nil, fmt.Errorf("failed to query feedback: %w", err)
	}
	defer rows.Close()

	var results []*models.FeedbackWithSkill
	for rows.Next() {
		var it models.FeedbackWithSkill
		if err := rows.Scan(&it.Description, &it.Priority, &it.Status, &it.SkillName); err != nil {
			if err == sql.ErrNoRows {
				break
			}
			r.logger.Error(ctx, "failed to scan feedback row", "error", err)
			return nil, fmt.Errorf("failed to scan feedback row: %w", err)
		}
		results = append(results, &it)
	}

	if err := rows.Err(); err != nil {
		r.logger.Error(ctx, "rows error after scanning feedback", "error", err)
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return results, nil
}
