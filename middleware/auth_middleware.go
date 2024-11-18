package middleware

import (
	"go_test/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
			ctx.Abort()
			return
		}
		token = strings.TrimPrefix(token, "Bearer ")
		valid, err := auth.ValidateJWT(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			ctx.Abort()
			return
		}

		ctx.Set("username",valid.Username)
		ctx.Next()
	}
}