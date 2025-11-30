package repository

import (
	"context"
	"database/sql"
	"fmt"
	"hack_change/internal/models"
	postgres "hack_change/pkg/db"
	"hack_change/pkg/logger"
	"strconv"

	sq "github.com/Masterminds/squirrel"
)

type UploadRepo struct {
	db     *postgres.DB
	logger logger.Logger
}

func NewUploadRepository(db *postgres.DB) *UploadRepo {
	return &UploadRepo{db: db, logger: logger.New("upload_repo")}
}

// CreateAssignment inserts assignment metadata and returns created model with assignment id
func (r *UploadRepo) CreateAssignment(ctx context.Context, a *models.Assignment) (*models.Assignment, error) {
	// normalize/validate task_id: assignments.task_id is integer FK -> parse and verify existence
	tid, err := strconv.ParseInt(a.TaskID, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid task id: %w", err)
	}

	// check that referenced task exists
	var exists bool
	err = r.db.Db.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM tasks WHERE task_id = $1)", tid).Scan(&exists)
	if err != nil {
		r.logger.Error(ctx, "failed to check task existence", "error", err)
		return nil, fmt.Errorf("failed to validate task id: %w", err)
	}
	if !exists {
		return nil, fmt.Errorf("task not found: %d", tid)
	}

	query := sq.Insert("assignments").
		Columns("student_id", "task_id", "upload_date", "file_path", "status").
		Values(a.StudentID, tid, a.UploadDate, a.FilePath, a.Status).
		Suffix("RETURNING assignment_id").
		PlaceholderFormat(sq.Dollar)

	sqlStr, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build insert query: %w", err)
	}

	var insertedID int64
	if err := r.db.Db.QueryRowContext(ctx, sqlStr, args...).Scan(&insertedID); err != nil {
		if err == sql.ErrNoRows {
			// unlikely with RETURNING, but handle generically
			return nil, fmt.Errorf("failed to insert assignment: no id returned")
		}
		r.logger.Error(ctx, "failed to insert assignment", "error", err)
		return nil, fmt.Errorf("failed to insert assignment: %w", err)
	}

	a.ID = insertedID
	r.logger.Debug(ctx, "created assignment", "assignmentId", insertedID)
	return a, nil
}
