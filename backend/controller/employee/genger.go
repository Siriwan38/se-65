package controller

import (
	"net/http"

	"github.com/Siriwan38/se-65/entity"
	"github.com/gin-gonic/gin"
)

func GetGender(c *gin.Context) {
	var gender entity.Gender
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM genders WHERE id = ?", id).Scan(&gender).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": gender})
}

func ListGender(c *gin.Context) {
	var gender []entity.Gender
	if err := entity.DB().Raw("SELECT * FROM genders").Scan(&gender).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": gender})
}
