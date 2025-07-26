package middleware

import (
	"it-order-app/backend/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}

		// Format header harus "Bearer <token>"
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format, 'Bearer ' prefix missing"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token is missing"})
			return
		}

		token, err := jwtService.ValidateToken(tokenString)

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		// Jika token valid, ekstrak claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			return
		}

		userIDFloat, ok := claims["user_id"].(float64)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid user_id format in token"})
			return
		}
		userID := uint(userIDFloat)

		role, ok := claims["role"].(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid role format in token"})
			return
		}

		// Simpan informasi user di context untuk digunakan oleh handler selanjutnya
		c.Set("userID", userID)
		c.Set("role", role)

		// Lanjutkan ke handler berikutnya
		c.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")

		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "User role not found in context"})
			return
		}

		if role.(string) != "admin" {
			// Jika user bukan admin, tolak akses
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Access denied. Admin role required."})
			return
		}

		// Jika user adalah admin, lanjutkan ke handler
		c.Next()
	}
}
