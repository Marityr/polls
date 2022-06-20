package handler

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var identityKey = "id"

func HandlerInit() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"*"},
		AllowHeaders: []string{
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Access-Control-Allow-Origin",
			"X-CSRF-Token",
			"Authorization",
			"accept",
			"origin",
			"Cache-Control",
			"X-Requested-With",
		},
		ExposeHeaders: []string{
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Access-Control-Allow-Origin",
			"X-CSRF-Token",
			"Authorization",
			"accept",
			"origin",
			"Cache-Control",
			"X-Requested-With",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// users := r.Group("/auth")
	// {
	// 	users.POST("/signup", signUp)
	// }

	r.POST("/auth/signup", signUp)

	blockquiz := r.Group("/blockquiz")
	{
		blockquiz.GET("/all", BlockQuizAll)
		blockquiz.GET("/:id", BlockQuizId)
		blockquiz.POST("/add", BlockQuizAdd)
	}

	quiz := r.Group("/quiz")
	{
		quiz.GET("/all", QuizAll)
		quiz.GET("/:id", QuizId)
		quiz.POST("/add", QuizAdd)
	}

	questions := r.Group("/questions")
	{
		questions.GET("/all", QuestionsAll)
		questions.GET("/:id", QuestionsId)
		questions.POST("/add", QuestionsAdd)
	}

	answers := r.Group("/answers")
	{
		answers.GET("/all", AnswersAll)
		answers.GET("/:id", AnswersId)
		answers.POST("/add", AnswersAdd)
	}

	cause := r.Group("/cause")
	{
		cause.GET("/all", CauseAll)
		cause.GET("/:id", CauseId)
		cause.POST("/add", CauseAdd)
	}

	result := r.Group("/result")
	{
		result.GET("/all", ResultAll)
		result.GET("/:id", ResultId)
		result.POST("/add", ResultAdd)
	}

	// tmp := AuthMiddleware()

	// users.POST("/login", tmp.LoginHandler)

	// users.Use(tmp.MiddlewareFunc())
	// users.GET("/add", AddUser)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run()
}
