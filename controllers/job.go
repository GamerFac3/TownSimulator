package controllers

import (
	"net/http"

	"TownSimulator/models"

	"github.com/gin-gonic/gin"
)

func AllJobs(c *gin.Context) {
	var jobs []models.Job
	models.DB.Find(&jobs)

	c.JSON(http.StatusOK, gin.H{"data": jobs})
}

func Createjob(c *gin.Context) {
	var input models.CreateJobInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	job := models.Job{Name: input.Name, Description: input.Description, PersonID: input.PersonID, BuildingID: input.BuildingID}
	models.DB.Create(&job)

	c.JSON(http.StatusOK, gin.H{"data": job})
}

func Findjob(c *gin.Context) {
	var job models.Job

	if err := models.DB.Where("id = ?", c.Param("id")).First(&job).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": job})
}

func Updatejob(c *gin.Context) {
	var job models.Job
	if err := models.DB.Where("id = ?", c.Param("id")).First(&job).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var input models.UpdateJobInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Model(&job).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": job})
}
