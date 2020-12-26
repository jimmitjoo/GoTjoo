package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jimmitjoo/gotjoo/app/controllers"
)

func SetupRoutes() {
	router := gin.Default()

	// middleware - intercept requests to use our db controller
	router.Use(func(context *gin.Context) {
		// provide db variable to controllers
		//context.Set("db", db)
		context.Next()
	})

	Api(router)
	Web(router)

	router.NoRoute(controllers.NotFoundPage)
	router.Run("0.0.0.0:7500")
}