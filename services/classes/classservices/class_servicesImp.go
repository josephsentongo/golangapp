package classservices

import (
	"github.com/josephsentongo/golangapp/services/classes/classesrepository"
	"github.com/josephsentongo/golangapp/services/classes/modelclass"
)

type ClassServicesImp struct {
	repositoryClass classesrepository.ClassRepository
}

func NewClassServiceImp(repository classesrepository.ClassRepository) ClassesServices {
	return &ClassServicesImp{repositoryClass: repository}
}

// SaveClass implements ClassServices.
func (c *ClassServicesImp) SaveClass(class modelclass.Classes) error {
	err := c.repositoryClass.SaveClass(&class)

	if err != nil {
		return err
	}
	return nil

}

// AllClass implements ClassesServices.
func (c *ClassServicesImp) AllClass() []modelclass.Classes {

	result, err := c.repositoryClass.AllClass()
	if err != nil {
		return nil
	}
	return result
}
