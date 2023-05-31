package controllers

import (
	"github.com/gin-gonic/gin"
	"h-vistars/server/src/database"
	"h-vistars/server/src/models"
	"net/http"
)

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func Login(c *gin.Context) {

}
