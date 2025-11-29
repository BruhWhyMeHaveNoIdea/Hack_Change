package handler

import (
	"io"
	"net/http"

	"hack_change/internal/dto"
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
func (h *UploaderHandler) Upload(c *gin.Context) {
	uidAny, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusUnauthorized, dto.APIResponse{Error: &dto.APIError{Code: "unauthenticated", Message: "missing user id"}})
		return
	}
	studentID, ok := uidAny.(string)
	if !ok || studentID == "" {
		c.JSON(http.StatusUnauthorized, dto.APIResponse{Error: &dto.APIError{Code: "unauthenticated", Message: "invalid user id"}})
		return
	}

	taskID := c.PostForm("task_id")
	if taskID == "" {
		c.JSON(http.StatusBadRequest, dto.APIResponse{Error: &dto.APIError{Code: "invalid_request", Message: "missing task_id"}})
		return
	}

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.APIResponse{Error: &dto.APIError{Code: "invalid_request", Message: "file is required"}})
		return
	}
	defer file.Close()

	// read file content into memory and delegate saving to service
	content, err := io.ReadAll(file)
	if err != nil {
		h.logger.Error(c.Request.Context(), "failed to read uploaded file", "error", err)
		c.JSON(http.StatusInternalServerError, dto.APIResponse{Error: &dto.APIError{Code: "server_error", Message: "failed to read file"}})
		return
	}

	// persist via service: service will save bytes and create DB record
	resp, err := h.svc.UploadAssignment(c.Request.Context(), studentID, taskID, header.Filename, content)
	if err != nil {
		h.logger.Error(c.Request.Context(), "failed to persist assignment via service", "error", err)
		c.JSON(http.StatusInternalServerError, dto.APIResponse{Error: &dto.APIError{Code: "server_error", Message: "failed to persist assignment"}})
		return
	}

	c.JSON(http.StatusCreated, dto.APIResponse{Data: resp})
}
