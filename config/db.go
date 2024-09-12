package config

import (
	"errors"

	"github.com/josephsentongo/golangapp/services/classes/modelclass"
	"github.com/josephsentongo/golangapp/services/students/model"
	"github.com/josephsentongo/golangapp/services/subjects/subjectmodel"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnection() *gorm.DB {
	sqlInfo := "host=localhost user=postgres password=elvis123456 dbname=finalExam port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})

	if err != nil {
		errors.New("Database failed to connect")
		return nil
	}
	db.AutoMigrate(&modelclass.Classes{}, &model.Students{}, subjectmodel.Subjects{},&model.User{},&model.Posts{},model.Comments{},model.Movie{},model.Actor{})
	return db
}
