package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/FabianEspitia-it/workers-crud/handlers"
)

// Leaders routes
func LeaderRouter(r *gin.Engine) {
	router := r.Group("/")

	router.POST("/leaders", handlers.AddLeader)
	router.GET("/leaders", handlers.GetLeaders)
	router.GET("/leaders/:id", handlers.GetLeader)
	router.PUT("/leaders/:id", handlers.UpdateLeader)
	router.DELETE("/leaders/:id", handlers.DeleteLeader)
}
