package handler

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Marityr/polls.git/shema"
	"github.com/Marityr/polls.git/tools"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserName  string
	FirstName string
	LastName  string
}

// @Summary     SignIn
// @Tags        Auth
// @Description login
// @Accept      json
// @Produce     json
// @Param       input body shema.Login true "credentials"
// @Router      /login [post]
func Authenticator(c *gin.Context) (interface{}, error) {
	type login struct {
		Username string `form:"username" json:"username" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}
	var loginVals login
	if err := c.ShouldBind(&loginVals); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	userID := loginVals.Username
	password := loginVals.Password

	if (userID == "admin" && password == "admin") || (userID == "string" && password == "string") {
		return &User{
			UserName:  userID,
			LastName:  "Bo-Yi",
			FirstName: "Wu",
		}, nil
	}

	return nil, jwt.ErrFailedAuthentication
}

func HashPassword(password string) string {
	b := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(b, 8)

	if err != nil {
		fmt.Println(err)
	}

	return string(hash)
}

type signup struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// @Summary     SignUp
// @Tags        Auth
// @Description SignUp
// @Accept      json
// @Produce     json
// @Param       input body signup true "credentials"
// @Router      /signup [get]
func CreateUser(c *gin.Context) {
	var cnt shema.User
	var r response

	jsonErr := c.BindJSON(&cnt)

	cnt.Username = strings.ToLower(cnt.Username)
	cnt.Password = HashPassword(cnt.Password)

	if jsonErr != nil {
		r.Errors = append(r.Errors, jsonErr.Error())
		c.JSON(http.StatusBadRequest, &r)
		return
	}

	dbErr := tools.DB.Save(&cnt).Error
	if dbErr != nil {
		r.Errors = append(r.Errors, dbErr.Error())
		c.JSON(http.StatusBadRequest, &r)
	} else {
		r.Data = cnt
		c.JSON(http.StatusOK, "User Create OK")
	}
}

func AuthMiddleware() *jwt.GinJWTMiddleware {
	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		//TODO проверка токена пользователя
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &User{
				UserName: claims[identityKey].(string),
			}
		},
		Authenticator: Authenticator,
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*User); ok && v.UserName == "admin" {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	if err != nil {
		log.Println("JWT Error:" + err.Error())
	}

	return authMiddleware
}
