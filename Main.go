package main

import (
	"TownSimulator/controllers"
	"TownSimulator/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase() // new

	r.GET("/persons", controllers.AllPersons)
	r.POST("/persons", controllers.CreatePerson)
	r.GET("/persons/:id", controllers.FindPerson)
	r.PATCH("/persons/:id", controllers.UpdatePerson)

	r.Run()
}
