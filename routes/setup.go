package routes

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/jimmitjoo/gotjoo/app/controllers"
)

var ginLambda *ginadapter.GinLambda

func SetupRoutes(router *gin.Engine) {

	Api(router)
	Web(router)
	//Auth(router)

	router.NoRoute(controllers.NotFoundPage)

	ginLambda = ginadapter.New(router)

	//router.Run("0.0.0.0:7500")
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	req.Headers["Content-Type"] = "text/html"
	return ginLambda.ProxyWithContext(ctx, req)
}