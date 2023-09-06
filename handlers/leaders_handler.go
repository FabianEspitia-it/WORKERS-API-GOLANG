package handlers

import (
	"github.com/FabianEspitia-it/workers-crud/db"
	"github.com/FabianEspitia-it/workers-crud/models"
	"github.com/FabianEspitia-it/workers-crud/utils"
	"github.com/gin-gonic/gin"
)

// Add a new lider into db
func AddLeader(context *gin.Context) {
	var Data struct {
		LeaderName string `json:"leader_name" binding:"required"`
		Email      string `json:"email" binding:"required,email"`
		Password   string `json:"password" binding:"required"`
		CountryID  uint   `json:"country_id" binding:"required"`
	}

	if err := context.ShouldBindJSON(&Data); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if len(Data.Password) < 8 {
		context.JSON(400, gin.H{"error": "Password must be at least 8 characters"})
		return
	}

	hashed_password, _ := utils.HashPassword(Data.Password)
	leader := models.Leader{
		LeaderName: Data.LeaderName,
		Email:      Data.Email,
		Password:   hashed_password,
		CountryID:  Data.CountryID,
	}

	result := db.DB.Create(&leader)

	if result.Error != nil {
		context.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}

	context.JSON(200, gin.H{
		"message": "Leader added successfully",
		"leader": leader,
	})
}

// Show all leaders 
func GetLeaders(context *gin.Context) {
	var leaders []models.Leader

	if err := db.DB.Find(&leaders).Error; err != nil {
		context.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(200, gin.H{
		"leaders": leaders,
	})
}

// Show a leader
func GetLeader(context *gin.Context) {
	id := context.Param("id")

	var leader models.Leader
	result := db.DB.First(&leader, id)

	if result.Error != nil {
		context.JSON(400, gin.H{"error": "leader not found"})
		return
	}
	db.DB.Model(&leader).Association("Country").Find(leader.CountryID)

	context.JSON(200, gin.H{"leader": leader})
}

// Update a leader
func UpdateLeader(context *gin.Context) {

	id := context.Param("id")

	var Data struct {
		LeaderName string
		Email      string
		Password   string
		CountryID  uint
	}

	context.Bind(&Data)

	var leader models.Leader
	db.DB.First(&leader, id)

	hashed_password, _ := utils.HashPassword(Data.Password)
	db.DB.Model(&leader).Updates(models.Leader{
		LeaderName: Data.LeaderName,
		Email:      Data.Email,
		Password:   hashed_password,
		CountryID:  Data.CountryID,
	})

	context.JSON(200, gin.H{
		"message": "Information updated successfully",
		"leader":  leader,
	})

}

// Delete a leader
func DeleteLeader(context *gin.Context) {
	id := context.Param("id")

	db.DB.Delete(&models.Leader{}, id)

	context.JSON(200, gin.H{
		"Message": "Leader eliminated correctly",
	})
}
