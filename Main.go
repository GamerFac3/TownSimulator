package main

import (
	"TownSimulator/controllers"
	"TownSimulator/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/person", controllers.AllPersons)
	r.POST("/person", controllers.CreatePerson)
	r.GET("/person/:id", controllers.FindPerson)
	r.PATCH("/person/:id", controllers.UpdatePerson)

	r.GET("/job", controllers.AllJobs)
	r.POST("/job", controllers.Createjob)
	r.GET("/job/:id", controllers.Findjob)
	r.PATCH("/job/:id", controllers.Updatejob)
	r.Run()
}
