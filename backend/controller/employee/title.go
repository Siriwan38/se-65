package controller

import (
	"net/http"

	"github.com/Siriwan38/se-65/entity"
	"github.com/gin-gonic/gin"
)

func GetTitle(c *gin.Context) {
	var title entity.Title
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM titles WHERE id = ?", id).Scan(&title).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": title})
}

func ListTitle(c *gin.Context) {
	var title []entity.Title
	if err := entity.DB().Raw("SELECT * FROM titles").Scan(&title).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": title})
}
