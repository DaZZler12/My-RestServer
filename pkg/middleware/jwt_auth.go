package middleware

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWTAuth(ctx *gin.Context) {
	// Get the token from the Authorization header
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization header is missing"})
		ctx.Abort()
		return
	}

	// Extract the token from the header
	tokenString := authHeader

	// Validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Provide the secret key used to sign the tokens
		return []byte("your-secret-key"), nil
	})
	fmt.Println("Abhirup-Token:", token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		ctx.Abort()
		return
	}

	// Check if the token is valid
	if token.Valid {
		ctx.Next()
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		ctx.Abort()
		return
	}
}
