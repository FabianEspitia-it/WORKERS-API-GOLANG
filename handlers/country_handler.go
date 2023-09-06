package handlers

import (
	"github.com/FabianEspitia-it/workers-crud/db"
	"github.com/FabianEspitia-it/workers-crud/models"
	"github.com/gin-gonic/gin"
)

// Show all countries
func GetCountries(context *gin.Context) {
	var countries []models.Country

	db.DB.Find(&countries)

	context.JSON(200, gin.H{
		"countries": countries,
	})
}

// Add new countries
func AddCountries(context *gin.Context) {
	var Data struct {
		CountryName string `json:"country_name" binding:"required"`
	}

	if err := context.ShouldBindJSON(&Data); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	country := models.Country{
		CountryName: Data.CountryName,
	}

	result := db.DB.Create(&country)

	if result.Error != nil {
		context.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}

	context.JSON(200, gin.H{
		"message": "Country added",
		"country": country,
	})
}
