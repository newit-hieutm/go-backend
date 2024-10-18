package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type HomeController struct {}

func NewHomeController() *HomeController {
	return &HomeController{}
}

func (home *HomeController) Welcome(c *gin.Context) {
	fmt.Println("Welcome.")
}