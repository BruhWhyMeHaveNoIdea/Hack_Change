package repository

import (
	"context"
	"hack_change/internal/models"
)

type UserRepository interface {
	Register(ctx context.Context, user *models.User) error
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	GetByID(ctx context.Context, id string) (*models.User, error)
}

type UploadRepository interface {
	CreateAssignment(ctx context.Context, a *models.Assignment) (*models.Assignment, error)
}
