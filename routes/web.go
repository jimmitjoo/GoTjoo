package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jimmitjoo/gotjoo/app/controllers"
)

func Web(router *gin.Engine) {
	router.GET("/", controller.StartPage)
	router.GET("/about", controller.AboutPage)
	router.GET("/contact", controller.ContactPage)
}