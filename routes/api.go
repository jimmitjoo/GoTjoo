package routes

import (
	"github.com/gin-gonic/gin"
	apicontroller "github.com/jimmitjoo/gotjoo/app/controllers/api"
)

func Api(router *gin.Engine) *gin.RouterGroup {
	api := router.Group("api")
	{
		api.GET("/", apicontroller.StartPage)
		api.GET("/about", apicontroller.AboutPage)
		api.GET("/contact", apicontroller.ContactPage)
	}

	return api
}
