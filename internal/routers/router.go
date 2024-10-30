package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/newit-hieutm/go-backend/internal/controllers"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	homeController := controllers.NewHomeController()
	userController := controllers.NewUserController()

	v1 := r.Group("/v1")
	{
		v1.GET("/home", homeController.Welcome)
	}

	userRoutes := r.Group("users")
	{
		userRoutes.GET("myinfo", userController.MyInfo)
	}

	return r
}

func Pongfunc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}