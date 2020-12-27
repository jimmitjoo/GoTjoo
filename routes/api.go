package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jimmitjoo/gotjoo/app/controllers"
)

func Api(router *gin.Engine) *gin.RouterGroup {
	api := router.Group("api")
	{
		api.GET("/", controller.StartPage)
		api.GET("/about", controller.AboutPage)
		api.GET("/contact", controller.ContactPage)
	}

	return api
}
