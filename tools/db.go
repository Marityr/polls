package tools

import (
	"log"

	"github.com/Marityr/polls.git/shema"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dbURL := "postgres://postgres:1111@localhost:5432/polls"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		// отключение зависимости связей на уровне базы
		DisableForeignKeyConstraintWhenMigrating: true,
		// кеширование запроса
		PrepareStmt: true,
	})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(
		&shema.User{},
		&shema.Group{},
		&shema.Permissions{},
		&shema.Models{},
		// &shema.UserData{},

		&shema.BlockQuiz{},
		&shema.Quiz{},
		&shema.Questions{},
		&shema.Answers{},

		&shema.Cause{},

		&shema.Result{},
	)

	// Create
	// db.Create(&shema.Book{Id: 1, Title: "War and World", Author: "Tolstoy", Desc: "Author"})

	DB = db
}
