package main

import(
	"github.com/FabianEspitia-it/workers-crud/db"
	"github.com/FabianEspitia-it/workers-crud/models"

)


func init()  {
	db.Db_connection()
	
}

func main(){
	db.DB.AutoMigrate(&models.Worker{},&models.Position{}, &models.Country{}, &models.Leader{})
}