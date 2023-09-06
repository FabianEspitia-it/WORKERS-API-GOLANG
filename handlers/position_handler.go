package handlers

import (
	"github.com/FabianEspitia-it/workers-crud/db"
	"github.com/FabianEspitia-it/workers-crud/models"
	"github.com/gin-gonic/gin"
)

func GetPositions(context *gin.Context) {
	var positions []models.Position

	db.DB.Find(&positions)

	context.JSON(200, gin.H{
		"positions": positions,
	})
}

func AddPosition(context *gin.Context) {
	position := models.Position{PositionName: "gamer", Description: "lalala"}

	db.DB.Create(&position)

	context.JSON(200, gin.H{
		"position": position,
	})
}
