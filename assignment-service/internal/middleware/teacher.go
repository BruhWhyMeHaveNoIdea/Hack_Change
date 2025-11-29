package middleware

import (
	"net/http"

	"hack_change/internal/repository"

	"github.com/gin-gonic/gin"
)

// RequireTeacher пропускает только пользователей с ролью teacher.
// Ожидает, что предыдущий AuthMiddleware положил в контекст `userID` (string).
func RequireTeacher(userRepo repository.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		uidAny, ok := c.Get("userID")
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing user id in context"})
			return
		}

		userID, ok := uidAny.(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid user id in context"})
			return
		}

		user, err := userRepo.GetByID(c.Request.Context(), userID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "access denied"})
			return
		}

		if !user.IsTeacher() {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "only teachers allowed"})
			return
		}

		// сохраняем полный объект пользователя в контексте для хэндлеров
		c.Set("user", user)
		c.Next()
	}
}
