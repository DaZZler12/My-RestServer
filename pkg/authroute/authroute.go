package authroute

import (
	"github.com/Dazzler/My-RestServer/pkg/handler"
	"github.com/gin-gonic/gin"
)

// SetupAuthRoutes initializes the authentication routes
func SetupAuthRoutes(r *gin.Engine, uhandler handler.UserController) {
	authGroup := r.Group("/auth")

	// Register the user-related routes under the authGroup
	authGroup.POST("/signup", uhandler.Signup)
	authGroup.POST("/login", uhandler.Login)
}
