package routes

import (
	"github.com/gin-gonic/gin"
	webcontroller "github.com/jimmitjoo/gotjoo/app/controllers/web"
)

func Web(router *gin.Engine) {
	router.GET("/", webcontroller.StartPage)
	router.GET("/about", webcontroller.AboutPage)
	router.GET("/contact", webcontroller.ContactPage)
}