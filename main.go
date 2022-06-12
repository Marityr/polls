package main

import (
	"github.com/Marityr/polls/docs"
	"github.com/Marityr/polls/handler"
	"github.com/Marityr/polls/tools"
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
