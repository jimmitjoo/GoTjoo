package apicontroller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func StartPage(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"title": "GoTjoo",
	})
}

func AboutPage(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"title": "About GoTjoo",
	})
}

func ContactPage(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"title": "Contact GoTjoo",
	})
}