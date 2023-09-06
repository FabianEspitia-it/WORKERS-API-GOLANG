package routes


import(
	"github.com/gin-gonic/gin"
	"github.com/FabianEspitia-it/workers-crud/models"
	"github.com/FabianEspitia-it/workers-crud/initializers"
)


func Add_worker(context *gin.Context)  {

	var data struct{
		Name string 
		Email string
		RolID uint 
		CountryID uint 
	}

	context.Bind(&data)

	worker := models.Worker{Name: data.Name, Email: data.Email, RolID: data.RolID, CountryID: data.CountryID}
	result := initializers.DB.Create(&worker)

	if result.Error != nil{
		context.Status(400)
		return 
	}

	context.JSON(200, gin.H{
		"worker": worker,
	})

}