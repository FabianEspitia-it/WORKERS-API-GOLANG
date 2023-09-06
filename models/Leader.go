package models


import(
	"gorm.io/gorm"
)

type Leader struct {
    gorm.Model
    LeaderName      string `gorm:"type:varchar(100);not null" json:"name"`
    Email     string `gorm:"type:varchar(100);not null;unique_index" json:"email"`
    Password  string `gorm:"type:varchar(100);not null;unique_index" json:"password"`
    Rol       string `gorm:"default:Leader"`
    CountryID uint   `gorm:"not null" json:"country_id"`
}
