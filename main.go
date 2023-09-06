package main 

import (
	"github.com/gin-gonic/gin"
	"github.com/FabianEspitia-it/workers-crud/initializers"
	"github.com/FabianEspitia-it/workers-crud/routes"
)


func init(){
	initializers.Db_connection()
}

func main()  {
	router := gin.Default()

	router.POST("/worker", routes.Add_worker)
	

	router.Run()
}

