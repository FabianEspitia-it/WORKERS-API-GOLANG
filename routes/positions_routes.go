package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/FabianEspitia-it/workers-crud/handlers"
)

// Positions Routes
func PositionsRouter(r *gin.Engine) {
	router := r.Group("/")

	router.GET("/positions", handlers.GetPositions)
	router.POST("/positions", handlers.AddPosition)

}