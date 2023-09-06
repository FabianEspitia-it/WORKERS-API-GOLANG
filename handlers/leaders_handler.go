package handlers

import (
	"github.com/FabianEspitia-it/workers-crud/db"
	"github.com/FabianEspitia-it/workers-crud/models"
	"github.com/FabianEspitia-it/workers-crud/utils"
	"github.com/gin-gonic/gin"
)

func AddWorker(context *gin.Context) {
	var Data struct {
		WorkerName string `json:"worker_name" binding:"required"`
		Email      string `json:"email" binding:"required,email"`
		Password   string `json:"password" binding:"required"`
		RolID      uint   `json:"rol_id" binding:"required"`
		CountryID  uint   `json:"country_id" binding:"required"`
		LeaderID   uint   `json:"leader_id"`
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
	worker := models.Worker{
		WorkerName: Data.WorkerName,
		Email:      Data.Email,
		Password:   hashed_password,
		RolID:      Data.RolID,
		CountryID:  Data.CountryID,
		LeaderID:   Data.LeaderID,
	}

	result := db.DB.Create(&worker)

	if result.Error != nil {
		context.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}

	context.JSON(200, gin.H{"worker": worker})
}

func GetWorkers(context *gin.Context) {
	var workers []models.Worker

	if err := db.DB.Find(&workers).Error; err != nil {
		context.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(200, gin.H{
		"workers": workers,
	})
}

func GetWorker(context *gin.Context) {
	id := context.Param("id")

	var worker models.Worker
	result := db.DB.First(&worker, id)

	if result.Error != nil {
		context.JSON(400, gin.H{"error": "worker not found"})
		return
	}

	db.DB.Model(&worker).Association("Position").Find(worker.RolID)
	db.DB.Model(&worker).Association("Country").Find(worker.CountryID)
	db.DB.Model(&worker).Association("Leader").Find(worker.LeaderID)

	context.JSON(200, gin.H{"worker": worker})
}

func UpdateWorker(context *gin.Context) {

	id := context.Param("id")

	var Data struct {
		WorkerName string
		Email      string
		Password   string
		RolID      uint
		CountryID  uint
		LeaderID   uint
	}

	context.Bind(&Data)

	var worker models.Worker
	db.DB.First(&worker, id)

	hashed_password, _ := utils.HashPassword(Data.Password)
	db.DB.Model(&worker).Updates(models.Worker{
		WorkerName: Data.WorkerName,
		Email:      Data.Email,
		Password:   hashed_password,
		RolID:      Data.RolID,
		CountryID:  Data.CountryID,
		LeaderID:   Data.LeaderID,
	})

	context.JSON(200, gin.H{
		"message": "Information updated successfully",
		"worker":  worker,
	})

}

func DeleteWorker(context *gin.Context) {
	id := context.Param("id")

	db.DB.Delete(&models.Worker{}, id)

	context.JSON(200, gin.H{
		"Message": "Worker eliminated correctly",
	})
}
