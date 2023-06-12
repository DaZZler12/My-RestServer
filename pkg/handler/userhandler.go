package handler

import (
	"net/http"

	models "github.com/Dazzler/My-RestServer/pkg/models"
	"github.com/Dazzler/My-RestServer/pkg/services/userservice"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService userservice.UserService
}

func NewUserController(userservice userservice.UserService) UserController {
	return UserController{
		UserService: userservice,
	}
}

// Signup godoc
// @Summary      User signup
// @Description  Create a new user account
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user     body     models.User  true  "User object"
// @Success      200  {object}  gin.H         "message: User created successfully"
// @Failure      400  {object}  gin.H         "message: Bad Request"
// @Failure      500  {object}  gin.H         "message: Internal Server Error"
// @Router       /auth/signup [post]
func (uc *UserController) Signup(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.UserService.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

// Login godoc
// @Summary      User login
// @Description  Authenticate user and generate JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user     body     models.User  true  "User object"
// @Success      200  {object}  gin.H         "token: JWT token"
// @Failure      400  {object}  gin.H         "message: Bad Request"
// @Failure      401  {object}  gin.H         "message: Unauthorized"
// @Failure      500  {object}  gin.H         "message: Internal Server Error"
// @Router       /auth/login [post]
func (uc *UserController) Login(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	token, err := uc.UserService.Login(&user)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
