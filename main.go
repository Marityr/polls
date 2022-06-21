package main

import (
	"github.com/Marityr/polls.git/docs"
	"github.com/Marityr/polls.git/handler"
	"github.com/Marityr/polls.git/tools"
)

func SwaggerInit() {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Polls API"
	docs.SwaggerInfo.Description = "Polls swagger api examples"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}
}

// @securityDefinitions.apikey ApiKeyAuth
// @in   header
// @name Authorization

func main() {
	SwaggerInit()
	tools.Init()

	handler.HandlerInit()
}
