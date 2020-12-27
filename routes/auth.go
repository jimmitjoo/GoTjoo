package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jimmitjoo/gotjoo/app/controllers"
	"github.com/jimmitjoo/gotjoo/app/middlewares"
)

func Auth(router *gin.Engine) {
	router.GET("/login", controller.LoginPage)
	router.POST("/login", controller.LoginAction)

	router.GET("/register", controller.RegisterPage)
	router.POST("/register", controller.RegisterAction)

	router.GET("/forgot-password", controller.ForgotPasswordPage)
	router.POST("/forgot-password", controller.ForgotPasswordAction)

	private := router.Group("/")
	private.Use(middleware.Authenticate)
	{
		private.GET("/me", controller.MePage)
	}
}