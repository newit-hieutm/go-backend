package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/newit-hieutm/go-backend/internal/services"
	"github.com/newit-hieutm/go-backend/pkg/loggers"
	responseData "github.com/newit-hieutm/go-backend/pkg/response"
	"go.uber.org/zap"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

func (uc *UserController) MyInfo(c *gin.Context) {
	myInfo := uc.userService.GetUserInfo()
	logger := loggers.InitLogger()
	logger.Error("myInfo", zap.String("name", myInfo))
	responseData.RenderSuccess(c, myInfo, http.StatusOK, "success")
}
