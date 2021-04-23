package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prasadadireddi/scytale-restapi/models"
)

type CreateRegInput struct {
	SpiffeID string `json:"spiffeid" binding:"required"`
	Selectors string `json:"selectors" binding:"required"`
}

type UpdateRegInput struct {
	SpiffeID string `json:"spiffeid"`
	Selectors string `json:"selectors"`
}

// GET /workloads
// Find all workloads
func GetWorkloads(c *gin.Context) {
	var workloads []models.Registration
	models.DB.Find(&workloads)

	c.JSON(http.StatusOK, gin.H{"data": workloads})
}

// GET /workloads/:spiffid
// Find a selector
func FindSelector(c *gin.Context) {
	// Get model if exist
	var workload models.Registration
	if err := models.DB.Where("id = ?", c.Param("id")).First(&workload).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": workload})
}

// POST /workloads
// Create new workload
func CreateWorkload(c *gin.Context) {
	// Validate input
	var input CreateRegInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create workload
	workload := models.Registration{SpiffeID: input.SpiffeID, Selectors: input.Selectors}
	models.DB.Create(&workload)

	c.JSON(http.StatusOK, gin.H{"data": workload})
}

// PATCH /workload/:id
// Update a workload
func UpdateWorkload(c *gin.Context) {
	// Get model if exist
	var workload models.Registration
	if err := models.DB.Where("id = ?", c.Param("id")).First(&workload).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateRegInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&workload).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": workload})
}

// DELETE /workloads/:id
// Delete a workload
func DeleteWorkload(c *gin.Context) {
	// Get model if exist
	var workload models.Registration
	if err := models.DB.Where("id = ?", c.Param("id")).First(&workload).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&workload)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
