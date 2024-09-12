package repositorysubject

import (
	"github.com/josephsentongo/golangapp/services/subjects/subjectmodel"
	"gorm.io/gorm"
)

type SubjectsImpl struct {
	Db *gorm.DB
}

func NewSubjectsRepository(db *gorm.DB) IrepositorySubject {
	return &SubjectsImpl{Db: db}
}

// SaveSubject implements IrepositorySubject.
func (s *SubjectsImpl) SaveSubject(subject subjectmodel.Subjects) error {
	err := s.Db.Create(&subject)

	if err != nil {
		return nil
	}
	return nil
}
