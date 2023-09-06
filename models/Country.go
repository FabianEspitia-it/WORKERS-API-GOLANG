package models


import(
	"gorm.io/gorm"
)

type Country struct{
	gorm.Model
	CountryName string `gorm:"type:varchar(30);not null; unique_index" json:"Name"`
}