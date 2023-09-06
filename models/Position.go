package models

import (
	"gorm.io/gorm"
)

// Position Table
type Position struct {
	ID        uint           `gorm:"primaryKey"`
	PositionName string `gorm:"type:varchar(50);not null;unique_index"`
	Description  string
  	DeletedAt gorm.DeletedAt `gorm:"index"`
}
