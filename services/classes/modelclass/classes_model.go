package modelclass

type Classes struct {
	Id    uint64 `gorm:"primaryKey;not null" json:"id"`
	Level string `gorm:"not null" json:"level" validate:"required"`
	Class string `gorm:"not null" json:"class" validate:"required"`
}
