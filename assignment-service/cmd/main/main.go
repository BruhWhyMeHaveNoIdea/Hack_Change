package main

import (
	"context"
	"fmt"
	"hack_change/internal/handler"
	repo "hack_change/internal/repository"
	"hack_change/internal/service"
	"hack_change/pkg/config"
	postgres "hack_change/pkg/db"
	"hack_change/pkg/logger"
	"log"
	"os"

	"hack_change/internal/auth"
	"hack_change/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// Инициализация логгера
	mainLogger := logger.New("hack_change")

	// Базовый контекст
	ctx := context.Background()

	r := gin.Default()

	// Загрузка конфигурации
	cfg, err := config.LoadConfig()
	if err != nil {
		mainLogger.Error(ctx, "error loading config", "error", err)
		os.Exit(1)
	}

	if cfg == nil {
		mainLogger.Error(ctx, "error loading config: config is nil")
		os.Exit(1)
	}

	// Инициализация базы данных
	db, err := postgres.New(cfg.Postgres)
	if err != nil {
		fmt.Println(err)
		mainLogger.Error(ctx, "failed to init postgres", "error", err)
		os.Exit(1)
	}

	// Инициализация JWT сервиса
	jwtService := auth.NewJWTService(cfg.JWT)

	// Инициализация репозиториев и сервисов
	userRepo := repo.NewAuthUserRepository(db)
	uploadRepo := repo.NewUploadRepository(db)
	service := service.NewService(
		userRepo,
		uploadRepo,
		jwtService,
	)

	// Публичные роуты (без JWT)
	authHandler := handler.NewAuthHandler(service)
	auth := r.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
		auth.POST("/register", authHandler.Register)
	}

	// Загрузка заданий — защищённый роут
	uploader := handler.NewUploaderHandler(service)
	// подключаем JWT middleware
	authMw := middleware.AuthMiddleware(jwtService)
	asg := r.Group("/")
	asg.Use(authMw)
	{
		asg.POST("/upload", uploader.Upload)
	}

	log.Println("Server started on :8080")
	r.Run(":8080")
}
