package models


import(
	"gorm.io/gorm"
)

type Country struct{
	gorm.Model
	Name string `gorm:"type:varchar(30);not null; unique_index" json:"Name"`
}