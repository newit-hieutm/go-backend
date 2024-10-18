package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/newit-hieutm/go-backend/internal/services"
	responseData "github.com/newit-hieutm/go-backend/pkg/response"
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
	responseData.RenderSuccess(c, myInfo, http.StatusOK, "success")
}