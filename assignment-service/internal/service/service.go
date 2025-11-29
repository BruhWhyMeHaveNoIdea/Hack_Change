package service

import (
	"context"
	"errors"
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
}

// NewService создает сервис
func NewService(userRepo repo.UserRepository, jwtService auth.JWTService) *Service {
	return &Service{
		userRepo:   userRepo,
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

	token, err := s.jwtService.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	loginResp := dto.LoginResponse{
		Token: token,
		User:  dto.UserToResponse(user),
	}
	return &loginResp, nil
}