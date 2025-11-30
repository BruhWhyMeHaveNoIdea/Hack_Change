package handler

import (
	"context"
	"net/http"
	"strconv"

	"hack_change/internal/dto"
	"hack_change/internal/service"

	"github.com/gin-gonic/gin"
)

type FeedbackHandler struct {
	svc *service.Service
}

func NewFeedbackHandler(svc *service.Service) *FeedbackHandler {
	return &FeedbackHandler{svc: svc}
}

// GetFeedback handles GET /:assignmentId/feedback
// @Summary Get feedback for assignment
// @Description Returns feedback tickets for a given assignment id
// @Tags feedback
// @Accept json
// @Produce json
// @Param assignmentId path int true "assignment id"
// @Success 200 {array} dto.FeedbackItem
// @Failure 400 {object} dto.APIResponse
// @Failure 401 {object} dto.APIResponse
// @Failure 500 {object} dto.APIResponse
// @Security bearerAuth
// @Router /{assignmentId}/feedback [get]
func (h *FeedbackHandler) GetFeedback(c *gin.Context) {
	aidStr := c.Param("assignmentId")
	if aidStr == "" {
		c.JSON(http.StatusBadRequest, dto.APIResponse{Error: &dto.APIError{Code: "invalid_request", Message: "assignmentId required"}})
		return
	}
	aid, err := strconv.ParseInt(aidStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.APIResponse{Error: &dto.APIError{Code: "invalid_request", Message: "assignmentId must be integer"}})
		return
	}

	ctx := context.Background()
	items, err := h.svc.GetFeedbackByAssignment(ctx, aid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.APIResponse{Error: &dto.APIError{Code: "server_error", Message: err.Error()}})
		return
	}

	// return object with status and data array
	resp := struct {
		Status int         `json:"status"`
		Data   interface{} `json:"data"`
	}{
		Status: http.StatusOK,
		Data:   items,
	}

	c.JSON(http.StatusOK, resp)
}
