package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginPage(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"data": "This is a login page."})
}

func RegisterPage(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"data": "This is a register page."})
}

func ForgotPasswordPage(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"data": "This is a forgot password page."})
}

func LoginAction(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"data": "Login action."})
}

func RegisterAction(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"data": "Register action."})
}

func ForgotPasswordAction(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"data": "Forgot password action."})
}

func MePage(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"data": "Me page."})
}