package controller

import (
	"net/http"

	"github.com/Siriwan38/se-65/entity"
	"github.com/gin-gonic/gin"
)

// GET /:id
func GetPatientType(c *gin.Context) {
	var statusmed entity.StatusMed
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM patient_types WHERE id = ?", id).Scan(&statusmed).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": statusmed})
}

// GET
func ListPatientTypes(c *gin.Context) {
	var patienttypes []entity.PatientType
	if err := entity.DB().Raw("SELECT * FROM patient_types").Scan(&patienttypes).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": patienttypes})
}
