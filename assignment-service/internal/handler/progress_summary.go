package handler

import (
	"context"
	"net/http"
	"strconv"

	"hack_change/internal/dto"
	"hack_change/internal/service"

	"github.com/gin-gonic/gin"
)

type ProgressSummaryHandler struct {
	svc *service.Service
}

func NewProgressSummaryHandler(svc *service.Service) *ProgressSummaryHandler {
	return &ProgressSummaryHandler{svc: svc}
}

// GET /progress/summary?studentId=...&courseId=...
// @Summary Get progress summary
// @Description Returns total tasks and completed count for a student in a course
// @Tags progress
// @Accept json
// @Produce json
// @Param studentId query int true "student id"
// @Param courseId query int true "course id"
// @Success 200 {object} dto.ProgressSummaryResponse
// @Failure 400 {object} dto.APIResponse
// @Failure 401 {object} dto.APIResponse
// @Failure 500 {object} dto.APIResponse
// @Security bearerAuth
// @Router /progress/summary [get]
func (h *ProgressSummaryHandler) GetSummary(c *gin.Context) {
	sStr := c.Query("studentId")
	cStr := c.Query("courseId")
	if sStr == "" || cStr == "" {
		c.JSON(http.StatusBadRequest, dto.APIResponse{Error: &dto.APIError{Code: "invalid_request", Message: "studentId and courseId are required"}})
		return
	}
	sid, err := strconv.ParseInt(sStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.APIResponse{Error: &dto.APIError{Code: "invalid_request", Message: "studentId must be integer"}})
		return
	}
	cid, err := strconv.ParseInt(cStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.APIResponse{Error: &dto.APIError{Code: "invalid_request", Message: "courseId must be integer"}})
		return
	}

	ctx := context.Background()
	res, err := h.svc.GetProgressSummary(ctx, sid, cid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.APIResponse{Error: &dto.APIError{Code: "server_error", Message: err.Error()}})
		return
	}

	c.JSON(http.StatusOK, dto.APIResponse{Data: res})
}
