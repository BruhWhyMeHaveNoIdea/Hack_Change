package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"hack_change/internal/models"
	postgres "hack_change/pkg/db"
	"hack_change/pkg/logger"

	sq "github.com/Masterminds/squirrel"
)

type AuthUserRepository struct {
	db     *postgres.DB
	logger logger.Logger
}

func NewAuthUserRepository(db *postgres.DB) *AuthUserRepository {
	return &AuthUserRepository{db: db, logger: logger.New("repository")}
}

// Register inserts a new student into `students` table and sets returned student_id to user.ID
func (r *AuthUserRepository) Register(ctx context.Context, user *models.User) error {
	// build insert query (students.student_id is serial)
	query := sq.Insert("students").
		Columns("first_name", "last_name", "email", "password").
		Values(user.FirstName, user.LastName, user.Email, user.PasswordHash).
		Suffix("RETURNING student_id").
		PlaceholderFormat(sq.Dollar)

	sqlStr, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	var sid int64
	err = r.db.Db.QueryRowContext(ctx, sqlStr, args...).Scan(&sid)
	if err != nil {
		return fmt.Errorf("failed to insert student: %w", err)
	}

	user.ID = strconv.FormatInt(sid, 10)
	// since students table doesn't store role/created_at, we set defaults
	if user.Role == "" {
		user.Role = models.RoleStudent
	}
	r.logger.Debug(ctx, "student registered", "studentId", user.ID)
	return nil
}

func (r *AuthUserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	query := sq.Select("student_id", "first_name", "last_name", "email", "password").
		From("students").
		Where(sq.Eq{"email": email}).
		PlaceholderFormat(sq.Dollar)

	sqlStr, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	var sid sql.NullInt64
	var user models.User
	var pwd sql.NullString

	err = r.db.Db.QueryRowContext(ctx, sqlStr, args...).Scan(&sid, &user.FirstName, &user.LastName, &user.Email, &pwd)
	if err != nil {
		return nil, fmt.Errorf("failed to get student by email: %w", err)
	}

	if sid.Valid {
		user.ID = strconv.FormatInt(sid.Int64, 10)
	}
	if pwd.Valid {
		user.PasswordHash = pwd.String
	}
	// default role for students
	user.Role = models.RoleStudent

	r.logger.Debug(ctx, "got student by email", "studentId", user.ID)
	return &user, nil
}

func (r *AuthUserRepository) GetByID(ctx context.Context, id string) (*models.User, error) {
	// id is expected to be numeric string (student_id)
	sid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid id format: %w", err)
	}

	query := sq.Select("student_id", "first_name", "last_name", "email", "password").
		From("students").
		Where(sq.Eq{"student_id": sid}).
		PlaceholderFormat(sq.Dollar)

	sqlStr, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	var outSid int64
	var user models.User
	var pwd sql.NullString

	err = r.db.Db.QueryRowContext(ctx, sqlStr, args...).Scan(&outSid, &user.FirstName, &user.LastName, &user.Email, &pwd)
	if err != nil {
		return nil, fmt.Errorf("failed to get student by id: %w", err)
	}

	user.ID = strconv.FormatInt(outSid, 10)
	if pwd.Valid {
		user.PasswordHash = pwd.String
	}
	user.Role = models.RoleStudent

	r.logger.Debug(ctx, "got student by id", "studentId", user.ID)
	return &user, nil
}
