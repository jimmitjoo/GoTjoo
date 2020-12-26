package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jimmitjoo/gotjoo/app/controllers"
)

func Api(router *gin.Engine) *gin.RouterGroup {
	api := router.Group("api")
	{
		api.GET("/", controllers.StartPage)
		api.GET("/about", controllers.AboutPage)
		api.GET("/contact", controllers.ContactPage)
	}

	return api
}
