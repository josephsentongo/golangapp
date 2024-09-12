package repository

import (
	"net/http"

	"github.com/josephsentongo/golangapp/helper"
	"github.com/josephsentongo/golangapp/services/students/model"
)

type StudentRepository interface {
	SaveStudent(student *model.Students) (*model.Students, error)
	StudentByID(studentId uint64) (model.Students, error)
	UpdateStudent(student *model.Students, studentId uint64) (*model.Students, error)
	DeleteStudent(studentId uint64) error
	FindAll(req *http.Request) ([]model.Students, int64,helper.Pagination, error)
	SaveUser(user *model.User) ( error)
	FindUsers(page int, pageSize int) ([]model.User, int64, error)
	SavePost(posts *model.Posts) (error)
	FindPosts(page int, pageSize int) ([]model.Posts, int64, error)
	SaveComment(comment *model.Comments) (error)
	FindComments(page int, pageSize int) ([]model.Comments, int64, error)
}
