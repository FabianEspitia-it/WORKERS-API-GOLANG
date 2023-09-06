package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/FabianEspitia-it/workers-crud/handlers"
)

// Countries routes
func CountriesRouter(r *gin.Engine) {
	router := r.Group("/")

	router.GET("/countries", handlers.GetCountries)
	router.POST("/countries", handlers.AddCountries)

}