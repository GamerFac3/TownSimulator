package controllers

import (
	"net/http"

	"TownSimulator/models"

	"github.com/gin-gonic/gin"
)

func AllbuildingType(c *gin.Context) {
	var buildingTypes []models.BuildingType
	// get all buildingTypes with all their relations
	models.DB.Preload("JobNames").Preload("Buildings").Find(&buildingTypes)

	c.JSON(http.StatusOK, gin.H{"data": buildingTypes})
}

func CreatebuildingType(c *gin.Context) {
	var input models.CreateBuildingTypeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	buildingType := models.BuildingType{Name: input.Name}
	models.DB.Preload("JobNames").Preload("Buildings").Create(&buildingType)

	c.JSON(http.StatusOK, gin.H{"data": buildingType})
}

func FindbuildingType(c *gin.Context) {
	var buildingType models.BuildingType

	if err := models.DB.Preload("JobNames").Preload("Buildings").Where("id = ?", c.Param("id")).First(&buildingType).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": buildingType})
}

func UpdatebuildingType(c *gin.Context) {
	var buildingType models.BuildingType
	if err := models.DB.Preload("JobNames").Preload("Buildings").Where("id = ?", c.Param("id")).First(&buildingType).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var input models.UpdateBuildingTypeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Model(&buildingType).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": buildingType})
}
