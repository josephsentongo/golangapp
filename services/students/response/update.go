package response

type UpdateStudents struct {
	Id    uint   `json:"id" validate:"required"`
	Name  string `json:"name" validate:"required"`
	Age   uint8  `json:"age" validate:"required"`
	Class string `json:"class" validate:"required"`
}
