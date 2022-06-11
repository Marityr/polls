package tools

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dbURL := "postgres://mojas:1111@localhost:5432/gobase"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		// отключение зависимости связей на уровне базы
		DisableForeignKeyConstraintWhenMigrating: true,
		// кеширование запроса
		PrepareStmt: true,
	})
	if err != nil {
		log.Fatalln(err)
	}

	// db.AutoMigrate(
	// )

	// Create
	// db.Create(&shema.Book{Id: 1, Title: "War and World", Author: "Tolstoy", Desc: "Author"})

	DB = db
}
