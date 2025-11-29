package handler

import (
	"net/http"

	"hack_change/internal/dto"
	"hack_change/internal/service"
	"hack_change/pkg/logger"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	svc    service.AuthService
	logger logger.Logger
}

func NewAuthHandler(svc service.AuthService) *AuthHandler {
	return &AuthHandler{svc: svc, logger: logger.New("handlers")}
}

// Register принимает JSON -> вызывает сервис -> возвращает UserResponse
// @Summary Register new student
// @Description Register new student account
// @Tags auth
// @Accept json
// @Produce json
// @Param body body dto.RegisterRequest true "register request"
// @Success 201 {object} dto.UserResponse
// @Failure 400 {object} dto.APIResponse
// @Failure 500 {object} dto.APIResponse
// @Router /auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.APIResponse{Error: &dto.APIError{Code: "invalid_request", Message: err.Error()}})
		return
	}

	ctx := c.Request.Context()
	userResp, err := h.svc.Register(ctx, &req)
	if err != nil {
		h.logger.Warn(ctx, "register failed", "error", err.Error())
		// простая обработка ошибок; можно расширить
		c.JSON(http.StatusBadRequest, dto.APIResponse{Error: &dto.APIError{Code: "register_failed", Message: err.Error()}})
		return
	}

	c.JSON(http.StatusCreated, dto.APIResponse{Data: userResp})
}

// Login принимает учетные данные, возвращает token + user
// @Summary Login student
// @Description Login with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param body body dto.LoginRequest true "login request"
// @Success 200 {object} dto.LoginResponse
// @Failure 400 {object} dto.APIResponse
// @Failure 401 {object} dto.APIResponse
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.APIResponse{Error: &dto.APIError{Code: "invalid_request", Message: err.Error()}})
		return
	}

	ctx := c.Request.Context()
	resp, err := h.svc.Login(ctx, &req)
	if err != nil {
		h.logger.Warn(ctx, "login failed", "error", err.Error())
		c.JSON(http.StatusUnauthorized, dto.APIResponse{Error: &dto.APIError{Code: "auth_failed", Message: err.Error()}})
		return
	}

	c.JSON(http.StatusOK, dto.APIResponse{Data: resp})
}
