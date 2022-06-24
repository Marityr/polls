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
	rd := gin.Default()

	rd.Use(cors.New(cors.Config{
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
	tmp := AuthMiddleware()
	r := rd.Group("api/v1")
	r.GET("/signup", CreateUser)
	r.POST("/login", tmp.LoginHandler)

	//r.Use(tmp.MiddlewareFunc())

	blockquiz := r.Group("/blockquiz")
	{
		blockquiz.GET("/", BlockQuiz)
		blockquizOption := blockquiz.Group("")
		blockquizOption.Use(tmp.MiddlewareFunc())
		{
			blockquizOption.POST("/add", AddBlockQuiz)
			blockquizOption.PUT("/put", UpdateBlockQuiz)
			blockquizOption.DELETE("/delete", DeleteBlockQuiz)
		}
	}

	quiz := r.Group("/quiz")
	{
		quiz.GET("/", Quiz)
		quizOption := quiz.Group("")
		quizOption.Use(tmp.MiddlewareFunc())
		{
			quizOption.POST("/add", AddQuiz)
			quizOption.PUT("/put", UpdateQuiz)
			quizOption.DELETE("/delete", DeleteQuiz)
		}
	}

	questions := r.Group("/questions")
	{
		questions.GET("/", Questions)
		questionsOption := questions.Group("")
		questionsOption.Use(tmp.MiddlewareFunc())
		{
			questionsOption.POST("/add", AddQuestions)
			questionsOption.PUT("/put", UpdateQuestions)
			questionsOption.DELETE("/delete", DeleteQuestions)
		}
	}

	answers := r.Group("/answers")
	{
		answers.GET("/", Answers)
		answersOption := answers.Group("")
		answersOption.Use(tmp.MiddlewareFunc())
		{
			answersOption.POST("/add", AddAnswers)
			answersOption.PUT("/put", UpdateAnswers)
			answersOption.DELETE("/delete", DeleteAnswers)
		}
	}

	cause := r.Group("/cause")
	{
		cause.GET("/", Cause)
		causeOption := cause.Group("")
		causeOption.Use(tmp.MiddlewareFunc())
		{
			causeOption.POST("/add", AddCause)
			causeOption.PUT("/put", UpdateCause)
			causeOption.DELETE("/delete", DeleteCause)
		}
	}

	result := r.Group("/result")
	{
		result.GET("/", Result)
		resultOption := result.Group("")
		resultOption.Use(tmp.MiddlewareFunc())
		{
			resultOption.POST("/add", AddResult)
			resultOption.PUT("/put", UpdateResult)
			resultOption.DELETE("/delete", DeleteResult)
		}
	}

	rd.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.POST("/upload", UploadFile)

	rd.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	rd.Run()
}

//TODO отправка файлов на сервер
