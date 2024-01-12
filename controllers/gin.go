package controllers

import (
	"errors"
	"net/http"

	"github.com/Clavavin/clavavin-api/database"
	"github.com/gin-gonic/gin"
)

// CreateWine creates a wine entry in database
func CreateWine(c *gin.Context) {
	var wine *database.Wine
	err := c.ShouldBind(&wine)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	res := database.DB.Create(wine)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error creating a wine",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"wine": wine,
	})
}

// ReadWine retrieves a wine entry from database
func ReadWine(c *gin.Context) {
	var wine database.Wine
	id := c.Param("id")
	res := database.DB.Find(&wine, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "wine not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"wine": wine,
	})
}

// ReadWines retrieves all wine entries from database
func ReadWines(c *gin.Context) {
	var wines []database.Wine
	res := database.DB.Find(&wines)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("authors not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"wines": wines,
	})
}

// UpdateWine updates a wine entry in database
func UpdateWine(c *gin.Context) {
	var wine database.Wine
	id := c.Param("id")
	err := c.ShouldBind(&wine)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	var updateWine database.Wine
	res := database.DB.Model(&updateWine).Where("id = ?", id).Updates(wine)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "wine not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"wine": wine,
	})
}

// DeleteWine deletes a wine entry in database
func DeleteWine(c *gin.Context) {
	var wine database.Wine
	id := c.Param("id")
	res := database.DB.Find(&wine, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "wine not found",
		})
		return
	}
	database.DB.Delete(&wine)
	c.JSON(http.StatusOK, gin.H{
		"message": "wine deleted successfully",
	})
}
