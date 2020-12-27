package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authenticate(context *gin.Context) {
	context.Redirect(http.StatusFound, "/login")
}
