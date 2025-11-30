package main

import (
	"context"
	"fmt"
	"hack_change/internal/handler"
	repo "hack_change/internal/repository"
	svc "hack_change/internal/service"
	"hack_change/pkg/config"
	postgres "hack_change/pkg/db"
	"hack_change/pkg/logger"
	"log"
	"os"

	"hack_change/internal/auth"
	"hack_change/internal/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Инициализация логгера
	mainLogger := logger.New("hack_change")

	fmt.Println(auth.HashPassword("password123"))

	// Базовый контекст
	ctx := context.Background()

	r := gin.Default()
	// CORS: разрешаем запросы от Swagger UI (контейнер на host localhost:5000)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5000", "http://127.0.0.1:5000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

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
	progressRepo := repo.NewProgressRepository(db)
	feedbackRepo := repo.NewFeedbackRepository(db)
	svcService := svc.NewService(
		userRepo,
		uploadRepo,
		progressRepo,
		feedbackRepo,
		jwtService,
	)

	// подключаем JWT middleware
	authMw := middleware.AuthMiddleware(jwtService)
	// Публичные роуты (без JWT)
	authHandler := handler.NewAuthHandler(svcService)
	auth := r.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
		auth.POST("/register", authHandler.Register)
	}

	// Health endpoints (public)
	healthHandler := handler.NewHealthHandler(db)
	r.GET("/health", healthHandler.Health)
	r.GET("/health/detailed", healthHandler.Detailed)

	// Загрузка заданий — защищённый роут
	uploader := handler.NewUploaderHandler(svcService)
	asg := r.Group("/")
	asg.Use(authMw)
	{
		asg.POST("/upload", uploader.Upload)
		progressHandler := handler.NewProgressHandler(svcService)
		asg.POST("/progress", progressHandler.CreateProgress)
		progressSummaryHandler := handler.NewProgressSummaryHandler(svcService)
		asg.GET("/progress/summary", progressSummaryHandler.GetSummary)
		feedbackHandler := handler.NewFeedbackHandler(svcService)
		asg.GET("/:assignmentId/feedback", feedbackHandler.GetFeedback)
	}

	log.Println("Server started on :8080")
	r.Run(":8080")
}
