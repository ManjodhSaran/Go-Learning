package middlewares

import (
	"net/http"

	"eventapi.com/utils"
	"github.com/gin-gonic/gin"
)

func Auth(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No authorization header provided"})
		return
	}

	userID, email, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	context.Set("user_id", userID)
	context.Set("email", email)
	context.Next()
}
