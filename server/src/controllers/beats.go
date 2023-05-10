package controllers

import (
	"github.com/gin-gonic/gin"
	"h-vistars/server/src/database"
	"h-vistars/server/src/models"
	"net/http"
)

func GetBeats(c *gin.Context) {
	var beats []models.Beat
	database.DB.Find(&beats)
	c.JSON(http.StatusOK, gin.H{"data": beats})
}

func GetBeatByID(c *gin.Context) {
	var beat models.Beat

	if err := database.DB.Where("id = ?", c.Param("id")).First(&beat).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Beat no encontrado!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": beat})
}

func CreateBeat(c *gin.Context) {
	var beat models.Beat
	if err := c.ShouldBindJSON(&beat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&beat)
	c.JSON(http.StatusOK, gin.H{"data": beat})
}

func UpdateBeat(c *gin.Context) {
	var beat models.Beat
	if err := database.DB.Where("id = ?", c.Param("id")).First(&beat).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Beat no encontrado!"})
		return
	}

	if err := c.ShouldBindJSON(&beat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&beat)
	c.JSON(http.StatusOK, gin.H{"data": beat})
}

func DeleteBeat(c *gin.Context) {
	var beat models.Beat
	if err := database.DB.Where("id = ?", c.Param("id")).First(&beat).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Beat no encontrado!"})
		return
	}

	database.DB.Delete(&beat)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
