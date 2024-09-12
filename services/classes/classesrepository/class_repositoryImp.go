package classesrepository

import (
	"github.com/josephsentongo/golangapp/services/classes/modelclass"
	"gorm.io/gorm"
)

type ClassRepositoryImp struct {
	Db *gorm.DB
}

func NewClassRepositoryImp(Db *gorm.DB) ClassRepository {

	return &ClassRepositoryImp{Db: Db}
}

// SaveClass implements ClassRepository.
func (c *ClassRepositoryImp) SaveClass(class *modelclass.Classes) error {
	result := c.Db.Create(&class)
	if result.Error != nil {
		return nil
	}
	return nil
}

// AllClass implements ClassRepository.
func (c *ClassRepositoryImp) AllClass() ([]modelclass.Classes, error) {
	var class []modelclass.Classes
	err := c.Db.Find(&class).Error
	if err != nil {
		return nil, err
	}

	return class, nil
}
