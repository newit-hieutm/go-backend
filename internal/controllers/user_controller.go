package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/newit-hieutm/go-backend/internal/services"
)

type UserController struct{
	userService *services.UserService
}

func NewUserController() *UserController{
	return &UserController{
		userService: services.NewUserService(),
	}
}


func (uc *UserController) MyInfo(c *gin.Context) {
	myInfo :=  uc.userService.GetUserInfo()
	c.JSON(http.StatusOK, gin.H{"myInfo": myInfo, "status": http.StatusOK})
}