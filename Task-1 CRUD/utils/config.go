package utils

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"crud-movies-api/model"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&model.Movie{})
	return db
}
