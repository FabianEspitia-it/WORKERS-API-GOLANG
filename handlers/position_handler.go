package handlers

import (
	"github.com/FabianEspitia-it/workers-crud/db"
	"github.com/FabianEspitia-it/workers-crud/models"
	"github.com/gin-gonic/gin"
)

// Show all positions
func GetPositions(context *gin.Context) {
	var positions []models.Position

	db.DB.Find(&positions)

	context.JSON(200, gin.H{
		"positions": positions,
	})
}

// Add a new position into database
func AddPosition(context *gin.Context) {
	var Data struct {
		PositionName string `json:"position:name" binding:"required"`
		Description string `json:"description" binding:"required"`
	}

	if err := context.ShouldBindJSON(&Data); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	position := models.Position{
		PositionName: Data.PositionName,
		Description: Data.Description,
	}

	result := db.DB.Create(&position)

	if result.Error != nil {
		context.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}

	context.JSON(200, gin.H{
		"message": "position added successfully",
		"country": position,
	})
}
