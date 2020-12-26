package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jimmitjoo/gotjoo/app/controllers"
)

func Web(router *gin.Engine) *gin.Engine {
	router.GET("/", controllers.StartPage)
	router.GET("/about", controllers.AboutPage)
	router.GET("/contact", controllers.ContactPage)

	return router
}