package main

import(
	"github.com/FabianEspitia-it/workers-crud/initializers"
	"github.com/FabianEspitia-it/workers-crud/models"

)


func init()  {
	initializers.Db_connection()
	
}

func main(){
	initializers.DB.AutoMigrate(&models.Worker{},&models.Position{}, &models.Country{}, &models.Leader{})
}