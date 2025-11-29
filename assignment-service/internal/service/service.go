package service

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"hack_change/internal/auth"
	"hack_change/internal/dto"
	"hack_change/internal/models"
	repo "hack_change/internal/repository"
	"hack_change/pkg/logger"
)

// AuthService описывает поведение сервиса аутентификации
type AuthService interface {
	Register(ctx context.Context, req *dto.RegisterRequest) (*dto.UserResponse, error)
	Login(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error)
}

type Service struct {
	userRepo   repo.UserRepository
	jwtService auth.JWTService
	logger     logger.Logger
	uploadRepo repo.UploadRepository
}

// NewService создает сервис
func NewService(userRepo repo.UserRepository, uploadRepo repo.UploadRepository, jwtService auth.JWTService) *Service {
	return &Service{
		userRepo:   userRepo,
		uploadRepo: uploadRepo,
		jwtService: jwtService,
		logger:     logger.New("service"),
	}
}

// Register регистрирует пользователя
func (s *Service) Register(ctx context.Context, req *dto.RegisterRequest) (*dto.UserResponse, error) {
	if req == nil {
		return nil, errors.New("empty request")
	}

	if _, err := s.userRepo.GetByEmail(ctx, req.Email); err == nil {
		return nil, errors.New("email already exists")
	}

	hashed, err := auth.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		// ID will be set by repository after insert (student_id)
		Email:        req.Email,
		PasswordHash: hashed,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		MiddleName:   req.MiddleName,
		Role:         models.Role(req.Role),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := s.userRepo.Register(ctx, user); err != nil {
		return nil, err
	}

	userResp := dto.UserToResponse(user)
	return &userResp, nil
}

// Login аутентифицирует и возвращает JWT
func (s *Service) Login(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error) {
	if req == nil {
		return nil, errors.New("empty request")
	}

	user, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !auth.CheckPassword(req.Password, user.PasswordHash) {
		return nil, errors.New("invalid credentials")
	}

	// JWT expects numeric user id; parse stored user.ID (string) to int64
	var uidInt int64
	if user.ID != "" {
		var parseErr error
		uidInt, parseErr = strconv.ParseInt(user.ID, 10, 64)
		if parseErr != nil {
			return nil, errors.New("invalid user id format for token generation")
		}
	}
	token, err := s.jwtService.GenerateToken(uidInt)
	if err != nil {
		return nil, err
	}

	loginResp := dto.LoginResponse{
		Token: token,
		User:  dto.UserToResponse(user),
	}
	return &loginResp, nil
}

// UploadAssignment saves file content to disk and creates assignment metadata record
func (s *Service) UploadAssignment(ctx context.Context, studentID, taskID, originalFilename string, content []byte) (*dto.AssignmentResponse, error) {
	if len(content) == 0 {
		return nil, fmt.Errorf("empty file content")
	}

	uploadsDir := filepath.Join("uploads", studentID)
	if err := os.MkdirAll(uploadsDir, 0o755); err != nil {
		return nil, fmt.Errorf("failed to create upload dir: %w", err)
	}

	fname := fmt.Sprintf("%d_%s", time.Now().UnixNano(), filepath.Base(originalFilename))
	dstPath := filepath.Join(uploadsDir, fname)

	if err := os.WriteFile(dstPath, content, 0o644); err != nil {
		return nil, fmt.Errorf("failed to write file: %w", err)
	}

	filePath := dstPath
	a := &models.Assignment{
		StudentID:  studentID,
		TaskID:     taskID,
		UploadDate: time.Now(),
		FilePath:   &filePath,
		Status:     "На проверке",
	}

	created, err := s.uploadRepo.CreateAssignment(ctx, a)
	if err != nil {
		return nil, err
	}

	resp := dto.AssignmentToResponse(created)
	return &resp, nil
}
