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

type ProgressRepo struct {
	db     *postgres.DB
	logger logger.Logger
}

func NewProgressRepository(db *postgres.DB) *ProgressRepo {
	return &ProgressRepo{db: db, logger: logger.New("progress_repo")}
}

// CreateProgress inserts a progress row and returns created model with progress_id
func (r *ProgressRepo) CreateProgress(ctx context.Context, p *models.Progress) (*models.Progress, error) {
	// validate task_id and student_id exist
	tid := p.TaskID
	sid := p.StudentID

	var exists bool
	err := r.db.Db.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM tasks WHERE task_id = $1)", tid).Scan(&exists)
	if err != nil {
		r.logger.Error(ctx, "failed to check task existence", "error", err)
		return nil, fmt.Errorf("failed to validate task id: %w", err)
	}
	if !exists {
		return nil, fmt.Errorf("task not found: %d", tid)
	}

	err = r.db.Db.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM students WHERE student_id = $1)", sid).Scan(&exists)
	if err != nil {
		r.logger.Error(ctx, "failed to check student existence", "error", err)
		return nil, fmt.Errorf("failed to validate student id: %w", err)
	}
	if !exists {
		return nil, fmt.Errorf("student not found: %d", sid)
	}

	// Build insert query using squirrel
	query := sq.Insert("progress").
		Columns("student_id", "task_id", "status", "score", "completion_date").
		Values(sid, tid, p.Status, p.Score, p.CompletionDate).
		Suffix("RETURNING progress_id").
		PlaceholderFormat(sq.Dollar)

	sqlStr, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build insert query: %w", err)
	}

	var insertedID int64
	if err := r.db.Db.QueryRowContext(ctx, sqlStr, args...).Scan(&insertedID); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("failed to insert progress: no id returned")
		}
		r.logger.Error(ctx, "failed to insert progress", "error", err)
		return nil, fmt.Errorf("failed to insert progress: %w", err)
	}

	p.ID = insertedID
	r.logger.Debug(ctx, "created progress", "progressId", insertedID)
	return p, nil
}

func (r *ProgressRepo) GetSummary(ctx context.Context, studentID int64, courseID int64) (*models.ProgressSummary, error) {
	joinClause := fmt.Sprintf("progress p ON t.task_id = p.task_id AND p.student_id = %d", studentID)

	qb := sq.Select("COUNT(t.task_id) AS total", "COUNT(p.task_id) AS completed").
		From("tasks t").
		LeftJoin(joinClause).
		Join("modules m ON t.module_id = m.module_id").
		Where(sq.Eq{"m.course_id": courseID}).
		PlaceholderFormat(sq.Dollar)

	sqlStr, args, err := qb.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build progress summary query: %w", err)
	}

	var res models.ProgressSummary
	if err := r.db.Db.QueryRowContext(ctx, sqlStr, args...).Scan(&res.Total, &res.Completed); err != nil {
		if err == sql.ErrNoRows {
			return &res, nil
		}
		r.logger.Error(ctx, "failed to query progress summary", "error", err)
		return nil, fmt.Errorf("failed to query progress summary: %w", err)
	}

	return &res, nil
}
