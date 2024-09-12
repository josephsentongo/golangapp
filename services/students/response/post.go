package response

import (
	"github.com/josephsentongo/golangapp/services/classes/modelclass"
	"github.com/josephsentongo/golangapp/services/students/model"
)

type PostStudents struct {
	Id    uint64             `json:"id"`
	Name  string             `json:"name" validate:"required"`
	Age   uint8              `validate:"gte=0,lte=130,required" json:"age"`
	Class modelclass.Classes `json:"class" validate:"required"`
}

func (ps PostStudents) ToModel() model.Students {
	return model.Students{
		Name:  ps.Name,
		Age:   ps.Age,
		Class: ps.Class,
	}
}
