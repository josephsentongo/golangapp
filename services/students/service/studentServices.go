package service

import (
	"net/http"

	"github.com/josephsentongo/golangapp/helper"
	"github.com/josephsentongo/golangapp/services/students/model"
	"github.com/josephsentongo/golangapp/services/students/response"
)

type StudentServices interface {
	SaveStudents(student model.Students) (*response.PostStudents, error)
	StudentByID(studentId uint64) (response.PostStudents, error)
	UpdateStudent(student model.Students, studentId uint64) (*response.PostStudents, error)
	DeleteStudentByID(studentId uint64) error
	Students(r *http.Request) ([]response.PostStudents, int64,helper.Pagination, error)
	SaveUser(user *model.User) ( error)
	FindUsers(page int, pageSize int) ([]model.User, int64, error)
	SavePost(posts *model.Posts) (error)
	FindPosts(page int, pageSize int) ([]model.Posts, int64, error)
	SaveComment(comment *model.Comments) (error)
	FindComments(page int, pageSize int) ([]model.Comments, int64, error)
}
