package servicesubject

import "github.com/josephsentongo/golangapp/services/subjects/subjectmodel"

type IServiceSubject interface {
	SaveSubject(subject subjectmodel.Subjects) error
}
