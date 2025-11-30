package handler

import (
	"io"
	"net/http"
	"strconv"

	svc "hack_change/internal/service"
	"hack_change/pkg/logger"

	"github.com/gin-gonic/gin"
)

// UploaderHandler сохраняет файлы и метаданные в таблицу assignments
type UploaderHandler struct {
	svc    *svc.Service
	logger logger.Logger
}

func NewUploaderHandler(svcObj *svc.Service) *UploaderHandler {
	return &UploaderHandler{svc: svcObj, logger: logger.New("uploader")}
}

// Upload принимает multipart/form-data с полями:
// - file: файл
// - task_id: id задания (string number)
// Авторизованный userID должен быть в контексте как "userID" (middleware.AuthMiddleware)
// @Summary Upload assignment file
// @Description Upload a file for an assignment (creates assignment record)
// @Tags assignments
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "file to upload"
// @Param task_id formData int true "task id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} dto.APIResponse
// @Failure 401 {object} dto.APIResponse
// @Failure 500 {object} dto.APIResponse
// @Security bearerAuth
// @Router /upload [post]
func (h *UploaderHandler) Upload(c *gin.Context) {
	uidAny, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "assignmentId": "0"})
		return
	}
	studentID, ok := uidAny.(string)
	if !ok || studentID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "assignmentId": "0"})
		return
	}

	taskID := c.PostForm("task_id")
	if taskID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "assignmentId": "0"})
		return
	}

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "assignmentId": "0"})
		return
	}
	defer file.Close()

	// read file content into memory and delegate saving to service
	content, err := io.ReadAll(file)
	if err != nil {
		h.logger.Error(c.Request.Context(), "failed to read uploaded file", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "assignmentId": "0"})
		return
	}

	// persist via service: service will save bytes and create DB record
	resp, err := h.svc.UploadAssignment(c.Request.Context(), studentID, taskID, header.Filename, content)
	if err != nil {
		h.logger.Error(c.Request.Context(), "failed to persist assignment via service", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "assignmentId": "0"})
		return
	}

	// parse assignment id to int64 to fetch feedback
	aid, err := strconv.ParseInt(resp.AssignmentID, 10, 64)
	if err != nil {
		// if parsing fails, still return OK with empty feedback
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "assignmentId": resp.AssignmentID, "feedback": []interface{}{}})
		return
	}

	// fetch feedback items for the assignment
	feedbackItems, err := h.svc.GetFeedbackByAssignment(c.Request.Context(), aid)
	if err != nil {
		// on error fetching feedback, return empty feedback list but still success
		h.logger.Warn(c.Request.Context(), "failed to fetch feedback for assignment", "error", err)
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "assignmentId": resp.AssignmentID, "feedback": []interface{}{}})
		return
	}

	// map dto.FeedbackItem to expected shape {description, skill, priority, status}
	var feedbackOut []gin.H
	for _, it := range feedbackItems {
		feedbackOut = append(feedbackOut, gin.H{
			"description": it.Description,
			"skill":       it.SkillName,
			"priority":    it.Priority,
			"status":      it.Status,
		})
	}

	// return only array of feedback objects as requested: [{description, skill, priority, status}]
	c.JSON(http.StatusOK, feedbackOut)
}
