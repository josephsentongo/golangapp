package classservices

import (
	"github.com/josephsentongo/golangapp/services/classes/modelclass"
)

type ClassesServices interface {
	SaveClass(class modelclass.Classes) error
	AllClass() []modelclass.Classes
}
