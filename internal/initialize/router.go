package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/newit-hieutm/go-backend/internal/controllers"
	"github.com/newit-hieutm/go-backend/internal/middlewares"
)

func InitRouter() {
	r := gin.Default()
	r.Use(middlewares.UserZapLogger())

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

	r.Run(":8888")
}
