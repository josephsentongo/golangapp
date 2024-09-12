package classesrepository

import "github.com/josephsentongo/golangapp/services/classes/modelclass"

type ClassRepository interface {
	SaveClass(class *modelclass.Classes) error
	AllClass() ([]modelclass.Classes, error)
}
