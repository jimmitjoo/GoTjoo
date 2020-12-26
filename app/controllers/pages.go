package controllers

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
	context.HTML(http.StatusOK, "about.gohtml", gin.H{
		"title": "About GoTjoo",
	})
}

func ContactPage(context *gin.Context) {
	context.HTML(http.StatusOK, "contact.gohtml", gin.H{
		"title": "Contact GoTjoo",
	})
}

func NotFoundPage(context *gin.Context) {
	context.JSON(http.StatusNotFound, gin.H{"data": "Something went wrong, we cannot find this page."})
}