package handler

import (
	"log"
	"strconv"

	"github.com/Marityr/polls.git/shema"
	"github.com/Marityr/polls.git/tools"
	"github.com/gin-gonic/gin"
)

// @Summary BlockQuizAll
// @Tags    BlockQuiz
// @Accept  json
//@Router  /blockquiz/all [get]
func GetBlockQuizAll(c *gin.Context) {
	var allblock []shema.BlockQuiz

	tools.DB.Where(shema.BlockQuiz{}).Find(&allblock)

	c.JSON(200, &allblock)
}

// @Summary BlockQuizID
// @Tags    BlockQuiz
// @Accept  json
// @Accept       json
// @Param        id     path   int   true   "id"
//@Router  /blockquiz/{id} [get]
func GetBlockQuizId(c *gin.Context) {
	var block []shema.BlockQuiz

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}

	tools.DB.Where(shema.BlockQuiz{Id: id}).Find(&block)

	c.JSON(200, &block)
}
