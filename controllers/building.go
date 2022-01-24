package controllers

import (
	"net/http"

	"TownSimulator/models"

	"github.com/gin-gonic/gin"
)

func AllBuilding(c *gin.Context) {
	var buildings []models.Building
	// get all buildings with all their relations
	models.DB.Preload("Jobs").Find(&buildings)

	c.JSON(http.StatusOK, gin.H{"data": buildings})
}

func Createbuilding(c *gin.Context) {
	var input models.CreateBuildingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	building := models.Building{Name: input.Name, BuildingTypeID: input.BuildingTypeID}
	models.DB.Preload("Jobs").Create(&building)

	c.JSON(http.StatusOK, gin.H{"data": building})
}

func Findbuilding(c *gin.Context) {
	var building models.Building

	if err := models.DB.Preload("Jobs").Where("id = ?", c.Param("id")).First(&building).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": building})
}

func Updatebuilding(c *gin.Context) {
	var building models.Building
	if err := models.DB.Preload("Jobs").Where("id = ?", c.Param("id")).First(&building).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var input models.CreateBuildingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Model(&building).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": building})
}
