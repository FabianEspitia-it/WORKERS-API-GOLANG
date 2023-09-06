package models

import (
	"gorm.io/gorm"
)


// Worker table
type Worker struct {
	ID        uint    `gorm:"primaryKey"`
	WorkerName string `gorm:"type:varchar(100);not null"`
	Email      string `gorm:"type:varchar(100);not null;unique_index"`
	Password   string `gorm:"type:varchar(100);not null;unique_index"`
	RolID      uint   `gorm:"not null"`
	CountryID  uint   `gorm:"not null"`
	LeaderID   uint   `gorm:"not null"`
  	DeletedAt gorm.DeletedAt `gorm:"index"`
}
