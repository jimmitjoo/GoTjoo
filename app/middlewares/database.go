package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jimmitjoo/gotjoo/config"
)

func Database(context *gin.Context) {
	db := config.ConnectDatabase()
	context.Set("db", db)
	context.Next()
}