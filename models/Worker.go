package models


import(
	"gorm.io/gorm"
)

type Worker struct{
	gorm.Model
	Name string `gorm:"type:varchar(100);not null" json:"name"`
	Email string `gorm:"type:varchar(100);not null;unique_index" json:"email"`
	RolID uint `gorm:"not null" json:"rol_id"`
	CountryID uint `gorm:"not null" json:"country_id"`
}

