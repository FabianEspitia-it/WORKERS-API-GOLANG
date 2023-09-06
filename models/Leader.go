package models

import (
	"gorm.io/gorm"
)
type Leader struct {
	ID        uint           `gorm:"primaryKey"`
	LeaderName string `gorm:"type:varchar(100);not null"`
	Email      string `gorm:"type:varchar(100);not null;unique_index"`
	Password   string `gorm:"type:varchar(100);not null;unique_index"`
	Rol        string `gorm:"default:Leader"`
	CountryID  uint   `gorm:"not null"`
  	DeletedAt gorm.DeletedAt `gorm:"index"`
}
