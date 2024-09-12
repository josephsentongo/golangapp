package subjectmodel

type Subjects struct {
	Id    uint64 `gorm:"primaryKey;not null" json:"id"`
	Name  string `gorm:"not null" json:"name" validate:"required"`
	Level string `gorm:"not null" json:"level" validate:"required"`
}
