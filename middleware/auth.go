package middleware

import (
	"net/http"

	"example.com/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		// because this is a middleware we want to abort if there's an error
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorised"})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorised"})
		return
	}

	// add userId to context object
	context.Set("userId", userId)
	// pass context to next handler
	context.Next()
}
