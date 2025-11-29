package handler

import (
	"context"
	"net/http"

	"hack_change/internal/dto"
	"hack_change/internal/service"

	"github.com/gin-gonic/gin"
)

type ProgressHandler struct {
	svc *service.Service
}

func NewProgressHandler(svc *service.Service) *ProgressHandler {
	return &ProgressHandler{svc: svc}
}

// CreateProgress handles POST /progress
// @Summary Create progress record
// @Description Record progress for a student and a task
// @Tags progress
// @Accept json
// @Produce json
// @Param body body dto.ProgressRequest true "progress request"
// @Success 201 {object} dto.ProgressResponse
// @Failure 400 {object} dto.APIResponse
// @Failure 401 {object} dto.APIResponse
// @Failure 500 {object} dto.APIResponse
// @Security bearerAuth
// @Router /progress [post]
func (h *ProgressHandler) CreateProgress(c *gin.Context) {
	var req dto.ProgressRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.APIResponse{Error: &dto.APIError{Code: "invalid_request", Message: err.Error()}})
		return
	}

	ctx := context.Background()
	resp, err := h.svc.CreateProgress(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.APIResponse{Error: &dto.APIError{Code: "server_error", Message: err.Error()}})
		return
	}

	c.JSON(http.StatusCreated, dto.APIResponse{Data: resp})
}
