package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func StartPage(context *gin.Context) {
	context.HTML(http.StatusOK, "index.gohtml", gin.H{
		"title": "GoTjoo",
	})
}

func AboutPage(context *gin.Context) {
	context.HTML(http.StatusOK, "page.gohtml", gin.H{
		"title": "About GoTjoo",
	})
}

func ContactPage(context *gin.Context) {
	context.HTML(http.StatusOK, "page.gohtml", gin.H{
		"title": "Contact GoTjoo",
	})
}

func NotFoundPage(context *gin.Context) {
	context.HTML(http.StatusNotFound, "error.gohtml", gin.H{
		"title": "Page not found",
	})
}