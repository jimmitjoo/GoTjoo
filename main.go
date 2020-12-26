package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gin-gonic/gin"
	"github.com/jimmitjoo/gotjoo/routes"
)

func init() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	routes.SetupRoutes(router)
}

func main() {
	lambda.Start(routes.Handler)
}