package main

import (
	"gitlab.com/dmitriy.woronow/go-rest-api.git/docs"
	"gitlab.com/dmitriy.woronow/go-rest-api.git/handler"
	"gitlab.com/dmitriy.woronow/go-rest-api.git/tools"
)

func SwaggerInit() {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Polls API"
	docs.SwaggerInfo.Description = "Polls swagger api examples"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"}
}

func main() {
	SwaggerInit()

	tools.Init()

	handler.HandlerInit()
}
