package models


import(
	"gorm.io/gorm"
)


type Position struct{
	gorm.Model
	Name string `gorm:"varchar(50);not null;unique_index" json: "Name"`
	description string `json: "description"`
}

