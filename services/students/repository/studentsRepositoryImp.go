package repository

import (
	"errors"
	"net/http"

	"github.com/josephsentongo/golangapp/helper"
	"github.com/josephsentongo/golangapp/services/students/model"
	"gorm.io/gorm"
)

type StudentsRepositoryImp struct {
	Db *gorm.DB
}

// FindAll implements StudentRepository.


// FindComments implements StudentRepository.
func (s *StudentsRepositoryImp) FindComments(page int, pageSize int) ([]model.Comments, int64, error) {
	var comments []model.Comments
	var totalRecords int64

	s.Db.Model(&model.Comments{}).Count(&totalRecords)

	offset := (page - 1) * pageSize
	result := s.Db.Limit(pageSize).Offset(offset).Find(&comments)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return comments, totalRecords, nil
}

// FindPosts implements StudentRepository.
func (s *StudentsRepositoryImp) FindPosts(page int, pageSize int) ([]model.Posts, int64, error) {
	var posts []model.Posts
	var totalRecords int64

	s.Db.Model(&model.Posts{}).Count(&totalRecords)

	offset := (page - 1) * pageSize
	result := s.Db.Select("posts.id,posts.name,Comments.comment").Preload("User").Preload("Comments").Limit(pageSize).Offset(offset).Find(&posts)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return posts, totalRecords, nil
}

// FindUsers implements StudentRepository.
func (s *StudentsRepositoryImp) FindUsers(page int, pageSize int) ([]model.User, int64, error) {
	var users []model.User
	var totalRecords int64
	s.Db.Model(&model.User{}).Count(&totalRecords)
	offset := (page - 1) * pageSize
	result := s.Db.Limit(pageSize).Offset(offset).Find(&users)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return users, totalRecords, nil
}

// SaveComment implements StudentRepository.
func (s *StudentsRepositoryImp) SaveComment(comment *model.Comments) error {
	err := s.Db.Create(comment)
	if err != nil {
		return err.Error
	}
	return nil
}

// SavePost implements StudentRepository.
func (s *StudentsRepositoryImp) SavePost(posts *model.Posts) error {
	err := s.Db.Create(posts)
	if err != nil {
		return err.Error
	}
	return nil
}

// SaveUser implements StudentRepository.
func (s *StudentsRepositoryImp) SaveUser(user *model.User) error {
	err := s.Db.Create(user)
	if err != nil {
		return err.Error
	}
	return nil
}

func NewStudentRepositoryImp(Db *gorm.DB) StudentRepository {

	return &StudentsRepositoryImp{Db: Db}
}

// SaveStudent implements StudentRepository.
func (s *StudentsRepositoryImp) SaveStudent(student *model.Students) (*model.Students, error) {

	result := s.Db.Create(&student)
	if result.Error != nil {
		return nil, result.Error
	}
	return student, nil

}

// StudentByID implements StudentRepository.
func (s *StudentsRepositoryImp) StudentByID(studentId uint64) (model.Students, error) {
	var student model.Students
	studentData := s.Db.Where("id=?", studentId).Find(&student)

	if studentData.Error != nil {
		return student, nil
	}
	return student, nil

}

// UpdateStudent implements StudentRepository.
func (s *StudentsRepositoryImp) UpdateStudent(student *model.Students, studentId uint64) (*model.Students, error) {

	var existingStudent model.Students
	if err := s.Db.First(&existingStudent, studentId).Error; err != nil {
		return nil, err
	}
	// Update the student fields
	result := s.Db.Model(&existingStudent).Updates(student)
	if result.Error != nil {
		return nil, result.Error
	}

	return &existingStudent, nil

}

// DeleteStudent implements StudentRepository.
func (s *StudentsRepositoryImp) DeleteStudent(studentId uint64) error {
	var students model.Students
	result := s.Db.Where("id = ?", studentId).Delete(&students)
	if result.Error != nil {
		return errors.New("failed to delete record: ")

	}
	return nil
}



func (s *StudentsRepositoryImp) FindAll(req *http.Request) ([]model.Students, int64,helper.Pagination, error) {
	var Students []model.Students
	var totalRecords int64

	s.Db.Model(&model.Students{}).Count(&totalRecords)
 scope,pagination:= helper.Paginate(req)
	
	result := s.Db.Preload("Class").Scopes(scope).Find(&Students)

	if result.Error != nil {
		return nil, 0,pagination, result.Error
	}

	return Students, totalRecords, pagination, nil
}

// FindAll implements StudentRepository.
// func (s *StudentsRepositoryImp) FindAll(page int, pageSize int) ([]model.Students, int64, error) {
// 	var Students []model.Students
// 	var totalRecords int64

// 	s.Db.Model(&model.Students{}).Count(&totalRecords)

// 	offset := (page - 1) * pageSize
// 	result := s.Db.Preload("Class").Limit(pageSize).Offset(offset).Find(&Students)
// 	if result.Error != nil {
// 		return nil, 0, result.Error
// 	}

// 	return Students, totalRecords, nil
// }
