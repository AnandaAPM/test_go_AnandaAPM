package controller

import (
	"go_test/models"
	"go_test/service"
	"go_test/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Service *service.UserService
	Validator *utils.Validator
}

func UserControllerDI(service *service.UserService) *UserController {
	validator := utils.Valid()

	return &UserController{
		Service:   service,
		Validator: validator,
	}
}

func (ctr *UserController) RegisterUser(ctx *gin.Context){

	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctr.Validator.Validate.Struct(user);err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	

	if err := ctr.Service.RegisterUser(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	

	ctx.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})

}

func (ctr *UserController) LoginUser(ctx *gin.Context){
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token,err := ctr.Service.LoginUser(&user)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"token": token})
	
}