package service

import (
	"net/http"

	"github.com/josephsentongo/golangapp/helper"
	"github.com/josephsentongo/golangapp/services/students/model"
	"github.com/josephsentongo/golangapp/services/students/repository"
	"github.com/josephsentongo/golangapp/services/students/response"
)

type StudentServicesImp struct {
	StudentRepository repository.StudentRepository
}

// FindComments implements StudentServices.
func (s *StudentServicesImp) FindComments(page int, pageSize int) ([]model.Comments, int64, error) {
students, totalRecords, err := s.StudentRepository.FindComments(page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	var studentResponses []model.Comments
	for _, student := range students {
		studentResponse := model.Comments{
			Id:    student.Id,
			Comment:  student.Comment,
			
		}
		studentResponses = append(studentResponses, studentResponse)
	}

	return studentResponses, totalRecords, nil
}

// FindPosts implements StudentServices.
func (s *StudentServicesImp) FindPosts(page int, pageSize int) ([]model.Posts, int64, error) {
	students, totalRecords, err := s.StudentRepository.FindPosts(page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	var studentResponses []model.Posts
	for _, student := range students {
		studentResponse := model.Posts{
			Id:    student.Id,
			Name:  student.Name,
			User: student.User,
			Comments: student.Comments,
			
		}
		studentResponses = append(studentResponses, studentResponse)
	}

	return studentResponses, totalRecords, nil
}

// FindUsers implements StudentServices.
func (s *StudentServicesImp) FindUsers(page int, pageSize int) ([]model.User, int64, error) {

	students, totalRecords, err := s.StudentRepository.FindUsers(page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	var studentResponses []model.User
	for _, student := range students {
		studentResponse := model.User{
			Id:    student.Id,
			Name:  student.Name,
			
		}
		studentResponses = append(studentResponses, studentResponse)
	}

	return studentResponses, totalRecords, nil
}

// SaveComment implements StudentServices.
func (s *StudentServicesImp) SaveComment(comment *model.Comments) error {
	err := s.StudentRepository.SaveComment(comment)
	if err != nil {
		return  err
	}
	return  nil
}

// SavePost implements StudentServices.
func (s *StudentServicesImp) SavePost(posts *model.Posts) error {
	err := s.StudentRepository.SavePost(posts)
	if err != nil {
		return err
	}
	return nil
}

// SaveUser implements StudentServices.
func (s *StudentServicesImp) SaveUser(user *model.User) error {
err := s.StudentRepository.SaveUser(user)
	if err != nil {
		return err
	}
	return nil
}

func NewStudentRepositoryImp(repository repository.StudentRepository) StudentServices {
	return &StudentServicesImp{StudentRepository: repository}
}

// SaveStudents implements StudentServices.
func (s *StudentServicesImp) SaveStudents(student model.Students) (*response.PostStudents, error) {

	savedStudent, err := s.StudentRepository.SaveStudent(&student)
	if err != nil {
		return nil, err
	}
	return &response.PostStudents{
		Id: savedStudent.Id,
	}, nil

}

// StudentByID implements StudentServices.
func (s *StudentServicesImp) StudentByID(studentId uint64) (response.PostStudents, error) {
	studentData, _ := s.StudentRepository.StudentByID(studentId)
	return response.PostStudents{
		Id:    studentData.Id,
		Name:  studentData.Name,
		Age:   studentData.Age,
		Class: studentData.Class,
	}, nil

}

// UpdateStudent implements StudentServices.
func (s *StudentServicesImp) UpdateStudent(student model.Students, studentId uint64) (*response.PostStudents, error) {
	savedStudent, err := s.StudentRepository.UpdateStudent(&student, studentId)
	if err != nil {
		return nil, err
	}
	return &response.PostStudents{
		Id: savedStudent.Id,
	}, nil

}

// DeleteStudentByID implements StudentServices.
func (s *StudentServicesImp) DeleteStudentByID(studentId uint64) error {
	s.StudentRepository.DeleteStudent(studentId)

	return nil
}

// Students implements StudentServices.
func (s *StudentServicesImp) Students(r *http.Request) ([]response.PostStudents, int64,helper.Pagination, error) {
	
	students,totalRecords,pagination, err := s.StudentRepository.FindAll(r)
	if err != nil {
		return nil, 0,pagination, err
	}

	var studentResponses []response.PostStudents
	for _, student := range students {
		studentResponse := response.PostStudents{
			Id:    student.Id,
			Name:  student.Name,
			Age:   student.Age,
			Class: student.Class,
		}
		studentResponses = append(studentResponses, studentResponse)
	}

	return studentResponses, totalRecords, pagination, nil
}
