package main 

import (
	"github.com/gin-gonic/gin"
	"github.com/FabianEspitia-it/workers-crud/initializers"
)


func init(){
	initializers.Load_variables()
}

func main()  {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	

	router.Run()
}

