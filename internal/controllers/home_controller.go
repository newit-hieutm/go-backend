package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	responseData "github.com/newit-hieutm/go-backend/pkg/response"
)

type HomeController struct{}

func NewHomeController() *HomeController {
	return &HomeController{}
}

func (home *HomeController) Welcome(c *gin.Context) {
	responseData.RenderSuccess(c, "Welcome to Home Page- Golang", http.StatusOK, "success")
}
