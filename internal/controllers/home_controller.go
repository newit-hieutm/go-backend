package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/newit-hieutm/go-backend/pkg/loggers"
	responseData "github.com/newit-hieutm/go-backend/pkg/response"
	"go.uber.org/zap"
)

type HomeController struct{}

func NewHomeController() *HomeController {
	return &HomeController{}
}

func (home *HomeController) Welcome(c *gin.Context) {
	logger := loggers.InitLogger()
	logger.Debug("myInfo", zap.String("name", "myInfo"))

	responseData.RenderSuccess(c, "Welcome to Home Page- Golang", http.StatusOK, "success")
}
