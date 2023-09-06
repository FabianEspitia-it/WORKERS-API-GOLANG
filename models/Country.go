package models

import (
	
)

type Country struct {
	ID        uint           `gorm:"primaryKey"`
	CountryName string `gorm:"type:varchar(30);not null; unique_index"`
}
