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
