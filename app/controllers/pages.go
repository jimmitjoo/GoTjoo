package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func StartPage(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"data": "This is the starting page."})
}

func AboutPage(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"data": "This is the about page."})
}

func ContactPage(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"data": "This is the contact page."})
}

func NotFoundPage(context *gin.Context) {
	context.JSON(http.StatusNotFound, gin.H{"data": "Something went wrong, we cannot find this page."})
}