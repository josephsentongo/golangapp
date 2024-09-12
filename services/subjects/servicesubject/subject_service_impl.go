package servicesubject

import (
	"github.com/josephsentongo/golangapp/services/subjects/repositorysubject"
	"github.com/josephsentongo/golangapp/services/subjects/subjectmodel"
)

type ServiceSubject struct {
	repositorySubject repositorysubject.IrepositorySubject
}

func NewServiceSubject(repository repositorysubject.IrepositorySubject) IServiceSubject {
	return &ServiceSubject{repositorySubject: repository}
}

// SaveSubject implements IServiceSubject.
func (s *ServiceSubject) SaveSubject(subject subjectmodel.Subjects) error {
	err := s.repositorySubject.SaveSubject(subject)
	if err != nil {
		return err
	}
	return err
}
