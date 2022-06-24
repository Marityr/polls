package handler

import (
	"net/http"

	"github.com/Marityr/polls.git/shema"
	"github.com/Marityr/polls.git/tools"
	"github.com/gin-gonic/gin"
)

type response struct {
	Errors []string
	Data   interface{}
}

// @Summary  BlockQuiz
// @Description  Блоки опросов
// @Tags     BlockQuiz
// @Accept   json
// @Produce  json
// @Param    id  query  string  false  "ID"
// @Router   /blockquiz/ [get]
func BlockQuiz(c *gin.Context) {
	var cnt []shema.BlockQuiz
	var r response

	q := tools.DB
	id := c.Copy().DefaultQuery("id", "")
	if id != "" {
		q = q.Where("id = ?", id)
	}

	dbErr := q.
		Preload("Quiz").
		Find(&cnt).Error

	if dbErr != nil {
		r.Errors = append(r.Errors, dbErr.Error())
		c.JSON(http.StatusBadRequest, &r)
	} else {
		r.Data = cnt
		c.JSON(http.StatusOK, &r)
	}
}

// @Summary AddBlockQuiz
// @Security ApiKeyAuth
// @Tags    BlockQuiz
// @Accept  json
// @Param   value   body   string   true   "Body"
// @Router  /blockquiz/add [post]
func AddBlockQuiz(c *gin.Context) {
	var cnt shema.BlockQuiz
	var r response

	jsonErr := c.BindJSON(&cnt)

	if jsonErr != nil {
		r.Errors = append(r.Errors, jsonErr.Error())
		c.JSON(http.StatusBadRequest, &r)
		return
	}

	dbErr := tools.DB.Create(&cnt).Error
	if dbErr != nil {
		r.Errors = append(r.Errors, dbErr.Error())
		c.JSON(http.StatusBadRequest, &r)
	} else {
		r.Data = cnt
		c.JSON(http.StatusOK, &r)
	}

}

// @Summary UpdateBlockQuiz
// @Security ApiKeyAuth
// @Tags    BlockQuiz
// @Accept  json
// @Param   value   body   string   true   "Body"
// @Router  /blockquiz/put [put]
func UpdateBlockQuiz(c *gin.Context) {
	var cnt shema.BlockQuiz
	var r response

	jsonErr := c.BindJSON(&cnt)

	if jsonErr != nil {
		r.Errors = append(r.Errors, jsonErr.Error())
		c.JSON(http.StatusBadRequest, &r)
		return
	}

	dbErr := tools.DB.
		Save(&cnt).
		Error
	if dbErr != nil {
		r.Errors = append(r.Errors, dbErr.Error())
		c.JSON(http.StatusBadRequest, &r)
	} else {
		r.Data = cnt
		c.JSON(http.StatusOK, &r)
	}

}

// @Summary  BlockQuizDelete
// @Security ApiKeyAuth
// @Tags     BlockQuiz
// @Accept   json
// @Param    id  query  string  true  "ID"
// @Router   /blockquiz/delete [Delete]
func DeleteBlockQuiz(c *gin.Context) {
	var cnt []shema.BlockQuiz
	var r response

	q := tools.DB
	id := c.Copy().DefaultQuery("id", "")
	if id != "" {
		q = q.Where("id = ?", id)
	}

	dbErr := q.Delete(&cnt).Error

	if dbErr != nil {
		r.Errors = append(r.Errors, dbErr.Error())
		c.JSON(http.StatusBadRequest, &r)
	} else {
		r.Data = cnt
		c.JSON(http.StatusOK, &r)
	}

}

// @Summary  Quiz
// @Description  Опрос
// @Tags     Quiz
// @Accept   json
// @Produce  json
// @Param    id  query  string  false  "ID"
// @Router   /quiz/ [get]
func Quiz(c *gin.Context) {
	var cnt []shema.Quiz
	var r response

	q := tools.DB
	id := c.Copy().DefaultQuery("id", "")
	if id != "" {
		q = q.Where("id = ?", id)
	}

	dbErr := q.
		Preload("Questions").
		Find(&cnt).Error

	if dbErr != nil {
		r.Errors = append(r.Errors, dbErr.Error())
		c.JSON(http.StatusBadRequest, &r)
	} else {
		r.Data = cnt
		c.JSON(http.StatusOK, &r)
	}
}

// @Summary AddQuiz
// @Security ApiKeyAuth
// @Tags    Quiz
// @Accept  json
// @Param   value   body   string   true   "Body"
// @Router  /quiz/add [post]
func AddQuiz(c *gin.Context) {
	var cnt shema.Quiz
	var r response

	jsonErr := c.BindJSON(&cnt)

	if jsonErr != nil {
		r.Errors = append(r.Errors, jsonErr.Error())
		c.JSON(http.StatusBadRequest, &r)
		return
	}

	dbErr := tools.DB.Create(&cnt).Error
	if dbErr != nil {
		r.Errors = append(r.Errors, dbErr.Error())
		c.JSON(http.StatusBadRequest, &r)
	} else {
		r.Data = cnt
		c.JSON(http.StatusOK, &r)
	}

}

// @Summary UpdateQuiz
// @Security ApiKeyAuth
// @Tags    Quiz
// @Accept  json
// @Param   value   body   string   true   "Body"
// @Router  /quiz/put [put]
func UpdateQuiz(c *gin.Context) {
	var cnt shema.Quiz
	var r response

	jsonErr := c.BindJSON(&cnt)

	if jsonErr != nil {
		r.Errors = append(r.Errors, jsonErr.Error())
		c.JSON(http.StatusBadRequest, &r)
		return
	}

	dbErr := tools.DB.
		Save(&cnt).
		Error
	if dbErr != nil {
		r.Errors = append(r.Errors, dbErr.Error())
		c.JSON(http.StatusBadRequest, &r)
	} else {
		r.Data = cnt
		c.JSON(http.StatusOK, &r)
	}

}

// @Summary  QuizDelete
// @Security ApiKeyAuth
// @Tags     Quiz
// @Accept   json
// @Param    id  query  string  true  "ID"
// @Router   /quiz/delete [Delete]
func DeleteQuiz(c *gin.Context) {
	var cnt []shema.Quiz
	var r response

	q := tools.DB
	id := c.Copy().DefaultQuery("id", "")
	if id != "" {
		q = q.Where("id = ?", id)
	}

	dbErr := q.Delete(&cnt).Error

	if dbErr != nil {
		r.Errors = append(r.Errors, dbErr.Error())
		c.JSON(http.StatusBadRequest, &r)
	} else {
		r.Data = cnt
		c.JSON(http.StatusOK, &r)
	}

}

// @Summary  Questions
// @Description  Вопросы
// @Tags     Questions
// @Accept   json
// @Produce  json
// @Param    id  query  string  false  "ID"
// @Router   /questions/ [get]
func Questions(c *gin.Context) {
	var cnt []shema.Questions
	var r response

	q := tools.DB
	id := c.Copy().DefaultQuery("id", "")
	if id != "" {
		q = q.Where("id = ?", id)
	}

	dbErr := q.
		Preload("Answers").
		Find(&cnt).Error

	if dbErr != nil {
		r.Errors = append(r.Errors, dbErr.Error())
		c.JSON(http.StatusBadRequest, &r)
	} else {
		r.Data = cnt
		c.JSON(http.StatusOK, &r)
	}
}

// @Summary AddQuestions
// @Security ApiKeyAuth
// @Tags    Questions
// @Accept  json
// @Param   value   body   string   true   "Body"
// @Router  /questions/add [post]
func AddQuestions(c *gin.Context) {
	var cnt shema.Questions
	var r response

	jsonErr := c.BindJSON(&cnt)

	if jsonErr != nil {
		r.Errors = append(r.Errors, jsonErr.Error())
		c.JSON(http.StatusBadRequest, &r)
		return
	}

	dbErr := tools.DB.Create(&cnt).Error
	if dbErr != nil {
		r.Errors = append(r.Errors, dbErr.Error())
		c.JSON(http.StatusBadRequest, &r)
	} else {
		r.Data = cnt
		c.JSON(http.StatusOK, &r)
	}

}

// @Summary UpdateQuestions
// @Security ApiKeyAuth
// @Tags    Questions
// @Accept  json
// @Param   value   body   string   true   "Body"
// @Router  /questions/put [put]
func UpdateQuestions(c *gin.Context) {
	var cnt shema.Questions
	var r response

	jsonErr := c.BindJSON(&cnt)

	if jsonErr != nil {
		r.Errors = append(r.Errors, jsonErr.Error())
		c.JSON(http.StatusBadRequest, &r)
		return
	}

	dbErr := tools.DB.
		Save(&cnt).
		Error
	if dbErr != nil {
		r.Errors = append(r.Errors, dbErr.Error())
		c.JSON(http.StatusBadRequest, &r)
	} else {
		r.Data = cnt
		c.JSON(http.StatusOK, &r)
	}

}

// @Summary  QuestionsDelete
// @Security ApiKeyAuth
// @Tags     Questions
// @Accept   json
// @Param    id  query  string  true  "ID"
// @Router   /questions/delete [Delete]
func DeleteQuestions(c *gin.Context) {
	var cnt []shema.Questions
	var r response

	q := tools.DB
	id := c.Copy().DefaultQuery("id", "")
	if id != "" {
		q = q.Where("id = ?", id)
	}

	dbErr := q.Delete(&cnt).Error

	if dbErr != nil {
		r.Errors = append(r.Errors, dbErr.Error())
		c.JSON(http.StatusBadRequest, &r)
	} else {
		r.Data = cnt
		c.JSON(http.StatusOK, &r)
	}

}

// @Summary  Answers
// @Description  Вопросы
// @Tags     Answers
// @Accept   json
// @Produce  json
// @Param    id  query  string  false  "ID"
// @Router   /answers/ [get]
func Answers(c *gin.Context) {
	var cnt []shema.Answers
	var r response

	q := tools.DB
	id := c.Copy().DefaultQuery("id", "")
	if id != "" {
		q = q.Where("id = ?", id)
	}

	dbErr := q.
		Preload("Cause").
		Find(&cnt).Error

	if dbErr != nil {
		r.Errors = append(r.Errors, dbErr.Error())
		c.JSON(http.StatusBadRequest, &r)
	} else {
		r.Data = cnt
		c.JSON(http.StatusOK, &r)
	}
}

// @Summary AddAnswers
// @Security ApiKeyAuth
// @Tags    Answers
// @Accept  json
// @Param   value   body   string   true   "Body"
// @Router  /answers/add [post]
func AddAnswers(c *gin.Context) {
	var cnt shema.Answers
	var r response

	jsonErr := c.BindJSON(&cnt)

	if jsonErr != nil {
		r.Errors = append(r.Errors, jsonErr.Error())
		c.JSON(http.StatusBadRequest, &r)
		return
	}

	dbErr := tools.DB.Create(&cnt).Error
	if dbErr != nil {
		r.Errors = append(r.Errors, dbErr.Error())
		c.JSON(http.StatusBadRequest, &r)
	} else {
		r.Data = cnt
		c.JSON(http.StatusOK, &r)
	}

}

// @Summary UpdateAnswers
// @Security ApiKeyAuth
// @Tags    Answers
// @Accept  json
// @Param   value   body   string   true   "Body"
// @Router  /answers/put [put]
func UpdateAnswers(c *gin.Context) {
	var cnt shema.Answers
	var r response

	jsonErr := c.BindJSON(&cnt)

	if jsonErr != nil {
		r.Errors = append(r.Errors, jsonErr.Error())
		c.JSON(http.StatusBadRequest, &r)
		return
	}

	dbErr := tools.DB.
		Save(&cnt).
		Error
	if dbErr != nil {
		r.Errors = append(r.Errors, dbErr.Error())
		c.JSON(http.StatusBadRequest, &r)
	} else {
		r.Data = cnt
		c.JSON(http.StatusOK, &r)
	}

}

// @Summary  AnswersDelete
// @Security ApiKeyAuth
// @Tags     Answers
// @Accept   json
// @Param    id  query  string  true  "ID"
// @Router   /answers/delete [Delete]
func DeleteAnswers(c *gin.Context) {
	var cnt []shema.Answers
	var r response

	q := tools.DB
	id := c.Copy().DefaultQuery("id", "")
	if id != "" {
		q = q.Where("id = ?", id)
	}

	dbErr := q.Delete(&cnt).Error

	if dbErr != nil {
		r.Errors = append(r.Errors, dbErr.Error())
		c.JSON(http.StatusBadRequest, &r)
	} else {
		r.Data = cnt
		c.JSON(http.StatusOK, &r)
	}

}

// @Summary  Cause
// @Description  Вопросы
// @Tags     Cause
// @Accept   json
// @Produce  json
// @Param    id  query  string  false  "ID"
// @Router   /cause/ [get]
func Cause(c *gin.Context) {
	var cnt []shema.Cause
	var r response

	q := tools.DB
	id := c.Copy().DefaultQuery("id", "")
	if id != "" {
		q = q.Where("id = ?", id)
	}

	dbErr := q.Find(&cnt).Error

	if dbErr != nil {
		r.Errors = append(r.Errors, dbErr.Error())
		c.JSON(http.StatusBadRequest, &r)
	} else {
		r.Data = cnt
		c.JSON(http.StatusOK, &r)
	}
}

// @Summary AddCause
// @Security ApiKeyAuth
// @Tags    Cause
// @Accept  json
// @Param   value   body   string   true   "Body"
// @Router  /cause/add [post]
func AddCause(c *gin.Context) {
	var cnt shema.Cause
	var r response

	jsonErr := c.BindJSON(&cnt)

	if jsonErr != nil {
		r.Errors = append(r.Errors, jsonErr.Error())
		c.JSON(http.StatusBadRequest, &r)
		return
	}

	dbErr := tools.DB.Create(&cnt).Error
	if dbErr != nil {
		r.Errors = append(r.Errors, dbErr.Error())
		c.JSON(http.StatusBadRequest, &r)
	} else {
		r.Data = cnt
		c.JSON(http.StatusOK, &r)
	}

}

// @Summary UpdateCause
// @Security ApiKeyAuth
// @Tags    Cause
// @Accept  json
// @Param   value   body   string   true   "Body"
// @Router  /cause/put [put]
func UpdateCause(c *gin.Context) {
	var cnt shema.Cause
	var r response

	jsonErr := c.BindJSON(&cnt)

	if jsonErr != nil {
		r.Errors = append(r.Errors, jsonErr.Error())
		c.JSON(http.StatusBadRequest, &r)
		return
	}

	dbErr := tools.DB.
		Save(&cnt).
		Error
	if dbErr != nil {
		r.Errors = append(r.Errors, dbErr.Error())
		c.JSON(http.StatusBadRequest, &r)
	} else {
		r.Data = cnt
		c.JSON(http.StatusOK, &r)
	}

}

// @Summary  CauseDelete
// @Security ApiKeyAuth
// @Tags     Cause
// @Accept   json
// @Param    id  query  string  true  "ID"
// @Router   /cause/delete [Delete]
func DeleteCause(c *gin.Context) {
	var cnt []shema.Cause
	var r response

	q := tools.DB
	id := c.Copy().DefaultQuery("id", "")
	if id != "" {
		q = q.Where("id = ?", id)
	}

	dbErr := q.Delete(&cnt).Error

	if dbErr != nil {
		r.Errors = append(r.Errors, dbErr.Error())
		c.JSON(http.StatusBadRequest, &r)
	} else {
		r.Data = cnt
		c.JSON(http.StatusOK, &r)
	}

}

// @Summary  Result
// @Description  Вопросы
// @Tags     Result
// @Accept   json
// @Produce  json
// @Param    id  query  string  false  "ID"
// @Router   /result/ [get]
func Result(c *gin.Context) {
	var cnt []shema.Result
	var r response

	q := tools.DB
	id := c.Copy().DefaultQuery("user_id", "")
	if id != "" {
		q = q.Where("user_id = ?", id)
	}

	dbErr := q.Find(&cnt).Error

	if dbErr != nil {
		r.Errors = append(r.Errors, dbErr.Error())
		c.JSON(http.StatusBadRequest, &r)
	} else {
		r.Data = cnt
		c.JSON(http.StatusOK, &r)
	}
}

// @Summary AddResult
// @Security ApiKeyAuth
// @Tags    Result
// @Accept  json
// @Param   value   body   string   true   "Body"
// @Router  /result/add [post]
func AddResult(c *gin.Context) {
	var cnt shema.Result
	var r response

	jsonErr := c.BindJSON(&cnt)

	if jsonErr != nil {
		r.Errors = append(r.Errors, jsonErr.Error())
		c.JSON(http.StatusBadRequest, &r)
		return
	}

	dbErr := tools.DB.Create(&cnt).Error
	if dbErr != nil {
		r.Errors = append(r.Errors, dbErr.Error())
		c.JSON(http.StatusBadRequest, &r)
	} else {
		r.Data = cnt
		c.JSON(http.StatusOK, &r)
	}

}

// @Summary UpdateResult
// @Security ApiKeyAuth
// @Tags    Result
// @Accept  json
// @Param   value   body   string   true   "Body"
// @Router  /result/put [put]
func UpdateResult(c *gin.Context) {
	var cnt shema.Result
	var r response

	jsonErr := c.BindJSON(&cnt)

	if jsonErr != nil {
		r.Errors = append(r.Errors, jsonErr.Error())
		c.JSON(http.StatusBadRequest, &r)
		return
	}

	dbErr := tools.DB.
		Save(&cnt).
		Error
	if dbErr != nil {
		r.Errors = append(r.Errors, dbErr.Error())
		c.JSON(http.StatusBadRequest, &r)
	} else {
		r.Data = cnt
		c.JSON(http.StatusOK, &r)
	}

}

// @Summary  ResultDelete
// @Security ApiKeyAuth
// @Tags     Result
// @Accept   json
// @Param    id  query  string  true  "ID"
// @Router   /result/delete [Delete]
func DeleteResult(c *gin.Context) {
	var cnt []shema.Result
	var r response

	q := tools.DB
	id := c.Copy().DefaultQuery("id", "")
	if id != "" {
		q = q.Where("id = ?", id)
	}

	dbErr := q.Delete(&cnt).Error

	if dbErr != nil {
		r.Errors = append(r.Errors, dbErr.Error())
		c.JSON(http.StatusBadRequest, &r)
	} else {
		r.Data = cnt
		c.JSON(http.StatusOK, &r)
	}

}
