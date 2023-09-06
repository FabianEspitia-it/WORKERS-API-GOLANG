package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/FabianEspitia-it/workers-crud/handlers"
)

func WorkerRouter(r *gin.Engine) {
	router := r.Group("/")

	router.POST("/workers", handlers.AddWorker)
	router.GET("/workers", handlers.GetWorkers)
	router.GET("/workers/:id", handlers.GetWorker)
	router.PUT("/workers/:id", handlers.UpdateWorker)
	router.DELETE("/workers/:id", handlers.DeleteWorker)
}
