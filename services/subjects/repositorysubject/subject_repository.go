package repositorysubject

import "github.com/josephsentongo/golangapp/services/subjects/subjectmodel"

type IrepositorySubject interface {
	SaveSubject(subject subjectmodel.Subjects) error
}
