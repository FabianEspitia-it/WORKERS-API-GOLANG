package models


import(
	"gorm.io/gorm"
)


type Position struct {
    gorm.Model
    PositionName        string `gorm:"type:varchar(50);not null;unique_index" json:"Name"`
    Description string `json:"description"`
}

