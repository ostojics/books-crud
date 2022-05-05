package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ostojics/books-crud/helpers"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "No auth header found"})
		context.Abort()
		return
	}

	splitToken := strings.Split(token, "Bearer")

	if len(splitToken) != 2 {
		fmt.Println("Failed")
	}

	parsedTokenString := strings.TrimSpace(splitToken[1])

	err := helpers.ValidateToken(parsedTokenString)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	context.Next()
}