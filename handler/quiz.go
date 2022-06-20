package handler

import (
	"log"
	"strconv"

	"github.com/Marityr/polls.git/shema"
	"github.com/Marityr/polls.git/tools"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @Summary BlockQuizAll
// @Tags    BlockQuiz
// @Accept  json
//@Router  /blockquiz/all [get]
func BlockQuizAll(c *gin.Context) {
	var allblock []shema.BlockQuiz

	tools.DB.Preload("Quiz").Find(&allblock)

	c.JSON(200, &allblock)
}

// @Summary BlockQuizID
// @Tags    BlockQuiz
// @Accept  json
// @Param   id     path   int   true   "id"
//@Router  /blockquiz/{id} [get]
func BlockQuizId(c *gin.Context) {
	var block []shema.BlockQuiz

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Info(err)
		return
	}

	tools.DB.Where(shema.BlockQuiz{Id: id}).Find(&block)

	c.JSON(200, &block)
}

// @Summary BlockQuizAdd
// @Tags    BlockQuiz
// @Accept  json
// @Param   value   body   string   true   "Body"
// @Router  /blockquiz/add [post]
func BlockQuizAdd(c *gin.Context) {
	var cnt shema.BlockQuiz

	c.BindJSON(&cnt)

	r := tools.DB.Create(&cnt) //cnt.Save(tools.DB)
	if r.RowsAffected == 0 {
		log.Println("Not Create")
		c.JSON(400, cnt)
		return
	}

	c.JSON(200, cnt)
}

// @Summary QuizAll
// @Tags    Quiz
// @Accept  json
//@Router  /quiz/all [get]
func QuizAll(c *gin.Context) {
	var allquiz []shema.Quiz

	tools.DB.Find(&allquiz)

	c.JSON(200, &allquiz)
}

// @Summary QuizID
// @Tags    Quiz
// @Accept  json
// @Param   id     path   int   true   "id"
//@Router  /quiz/{id} [get]
func QuizId(c *gin.Context) {
	var quiz []shema.Quiz

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Info(err)
		return
	}

	tools.DB.Where(shema.BlockQuiz{Id: id}).Find(&quiz)

	c.JSON(200, &quiz)
}

// @Summary QuizAdd
// @Tags    Quiz
// @Accept  json
// @Param   value   body   string   true   "Body"
// @Router  /quiz/add [post]
func QuizAdd(c *gin.Context) {
	var cnt shema.Quiz

	c.BindJSON(&cnt)
	tools.DB.Create(&cnt)

	c.JSON(200, cnt)
}

// @Summary QuestionsAll
// @Tags    Questions
// @Accept  json
//@Router  /questions/all [get]
func QuestionsAll(c *gin.Context) {
	var allquestions []shema.Questions

	tools.DB.Find(&allquestions)

	c.JSON(200, &allquestions)
}

// @Summary QuestionsID
// @Tags    Questions
// @Accept  json
// @Param   id     path   int   true   "id"
//@Router  /questions/{id} [get]
func QuestionsId(c *gin.Context) {
	var questions []shema.Questions

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Info(err)
		return
	}

	tools.DB.Where(shema.Questions{Id: id}).Find(&questions)

	c.JSON(200, &questions)
}

// @Summary QuestionsAdd
// @Tags    Questions
// @Accept  json
// @Param   value   body   string   true   "Body"
// @Router  /questions/add [post]
func QuestionsAdd(c *gin.Context) {
	var cnt shema.Questions

	c.BindJSON(&cnt)
	tools.DB.Create(&cnt)

	c.JSON(200, cnt)
}

// @Summary AnswersAll
// @Tags    Answers
// @Accept  json
//@Router  /answers/all [get]
func AnswersAll(c *gin.Context) {
	var allanswers []shema.Answers

	tools.DB.Find(&allanswers)

	c.JSON(200, &allanswers)
}

// @Summary AnswersID
// @Tags    Answers
// @Accept  json
// @Param   id     path   int   true   "id"
//@Router  /answers/{id} [get]
func AnswersId(c *gin.Context) {
	var answers []shema.Answers

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Info(err)
		return
	}

	tools.DB.Where(shema.Answers{Id: id}).Find(&answers)

	c.JSON(200, &answers)
}

// @Summary AnswersAdd
// @Tags    Answers
// @Accept  json
// @Param   value   body   string   true   "Body"
// @Router  /answers/add [post]
func AnswersAdd(c *gin.Context) {
	var cnt shema.Answers

	c.BindJSON(&cnt)
	tools.DB.Create(&cnt)

	c.JSON(200, cnt)
}

// @Summary CauseAll
// @Tags    Cause
// @Accept  json
//@Router  /cause/all [get]
func CauseAll(c *gin.Context) {
	var allcause []shema.Cause

	tools.DB.Find(&allcause)

	c.JSON(200, &allcause)
}

// @Summary CauseID
// @Tags    Cause
// @Accept  json
// @Param   id     path   int   true   "id"
//@Router  /cause/{id} [get]
func CauseId(c *gin.Context) {
	var cause []shema.Cause

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Info(err)
		return
	}

	tools.DB.Where(shema.Answers{Id: id}).Find(&cause)

	c.JSON(200, &cause)
}

// @Summary CauseAdd
// @Tags    Cause
// @Accept  json
// @Param   value   body   string   true   "Body"
// @Router  /cause/add [post]
func CauseAdd(c *gin.Context) {
	var cnt shema.Cause

	c.BindJSON(&cnt)
	tools.DB.Create(&cnt)

	c.JSON(200, cnt)
}

// @Summary ResultAll
// @Tags    Result
// @Accept  json
//@Router  /result/all [get]
func ResultAll(c *gin.Context) {
	var allresult []shema.Result

	tools.DB.Find(&allresult)

	c.JSON(200, &allresult)
}

// @Summary ResultID
// @Tags    Result
// @Accept  json
// @Param   id     path   int   true   "id"
//@Router  /result/{id} [get]
func ResultId(c *gin.Context) {
	var result []shema.Result

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Info(err)
		return
	}

	tools.DB.Where(shema.Result{Id: id}).Find(&result)

	c.JSON(200, &result)
}

// @Summary ResultAdd
// @Tags    Result
// @Accept  json
// @Param   value   body   string   true   "Body"
// @Router  /result/add [post]
func ResultAdd(c *gin.Context) {
	var cnt shema.Result

	c.BindJSON(&cnt)
	tools.DB.Create(&cnt)

	c.JSON(200, cnt)
}
