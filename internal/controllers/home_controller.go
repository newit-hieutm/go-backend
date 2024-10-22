package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/newit-hieutm/go-backend/pkg/loggers"
	"go.uber.org/zap"
)

type HomeController struct{}

func NewHomeController() *HomeController {
	return &HomeController{}
}

func (home *HomeController) Welcome(c *gin.Context) {
	logger := loggers.InitLogger()
	logger.Debug("myInfo", zap.String("name", "myInfo"))
}
