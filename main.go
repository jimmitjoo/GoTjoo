package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gin-gonic/gin"
	middleware "github.com/jimmitjoo/gotjoo/app/middlewares"
	"github.com/jimmitjoo/gotjoo/config"
	"github.com/jimmitjoo/gotjoo/routes"
)

func init() {

	router := gin.Default()

	if config.UseDatabase {
		router.Use(middleware.Database)
	}

	router.LoadHTMLGlob("templates/*.gohtml")
	routes.SetupRoutes(router)
}

func main() {
	lambda.Start(routes.Handler)
}