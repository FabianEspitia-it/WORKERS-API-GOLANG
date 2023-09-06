package main

import (
	"github.com/FabianEspitia-it/workers-crud/db"
	"github.com/FabianEspitia-it/workers-crud/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	db.Db_connection()
}

func main() {
	router := gin.Default()

	routes.WorkerRouter(router)
	routes.PositionsRouter(router)
	routes.CountriesRouter(router)
	routes.LeaderRouter(router)

	router.Run()
}
