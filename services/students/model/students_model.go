package model

import "github.com/josephsentongo/golangapp/services/classes/modelclass"

type Students struct {
	Id      uint64 `gorm:"primaryKey;not null"`
	Name    string
	Age     uint8
	ClassID uint
	Class   modelclass.Classes `gorm:"foreignKey:ClassID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}



type User struct {
	Id      uint64 `gorm:"primaryKey;not null" json:"id"`
	Name string `json:"name,omitempty"`

}

type Posts struct {
Id uint64 `gorm:"primaryKey;not null"`
Name string `json:"name,omitempty"`
UserId *uint64  `gorm:"index;"  json:"user_id,omitempty"`
User   User  `gorm:"foreignKey:UserId" json:"users,omitempty"`
Comments []Comments `json:"comments,omitempty"`
}

type Comments struct {
	Id     uint64 `gorm:"primaryKey;" json:"id,omitempty"`
	PostsId uint64 `gorm:"index;"  json:"post_id,omitempty"`
	Comment  string`json:"comment,omitempty"`
	UserId uint64 `gorm:"index;"  json:"user_id,omitempty"`
	// User   User  `json:"user,omitempty"`
}


type Movie struct {
	Id    uint64 `gorm:"primaryKey;" json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Actors  []Actor `gorm:"many2many:filmography;" json:"actors"`
}


type Actor struct {
	Id   uint64 `gorm:"primaryKey;" json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}