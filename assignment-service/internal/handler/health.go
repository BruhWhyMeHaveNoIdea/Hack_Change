package handler

import (
	"context"
	"net/http"
	"os"
	"path/filepath"
	"time"

	pg "hack_change/pkg/db"
	"hack_change/pkg/logger"

	"github.com/gin-gonic/gin"
)

// HealthHandler provides simple and detailed health checks
type HealthHandler struct {
	db     *pg.DB
	logger logger.Logger
}

func NewHealthHandler(db *pg.DB) *HealthHandler {
	return &HealthHandler{db: db, logger: logger.New("health")}
}

// Health returns a minimal health response
func (h *HealthHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "healthy", "service": "assignment"})
}

// Detailed performs quick checks: DB ping and uploads dir writeability
func (h *HealthHandler) Detailed(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	status := map[string]string{
		"service":           "assignment",
		"database":          "unknown",
		"uploads_directory": "unknown",
	}

	// DB check
	if h.db != nil && h.db.Db != nil {
		if err := h.db.Db.PingContext(ctx); err != nil {
			h.logger.Warn(ctx, "db ping failed", "error", err)
			status["database"] = "unreachable"
		} else {
			status["database"] = "healthy"
		}
	} else {
		status["database"] = "not_configured"
	}

	// Uploads dir check: ensure we can write a temp file
	uploadsDir := "uploads"
	if err := os.MkdirAll(uploadsDir, 0o755); err != nil {
		h.logger.Warn(ctx, "failed to ensure uploads dir", "error", err)
		status["uploads_directory"] = "unavailable"
	} else {
		tmp := filepath.Join(uploadsDir, ".hc_tmp")
		data := []byte(time.Now().Format(time.RFC3339))
		if err := os.WriteFile(tmp, data, 0o644); err != nil {
			h.logger.Warn(ctx, "uploads dir not writable", "error", err)
			status["uploads_directory"] = "not_writable"
		} else {
			_ = os.Remove(tmp)
			status["uploads_directory"] = "writable"
		}
	}

	c.JSON(http.StatusOK, status)
}
