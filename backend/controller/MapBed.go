package controller

import (
	"github.com/watanprasai/sa-65-example/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

// POST /users

func CreateMapBed(c *gin.Context) {

	var mapbed entity.Map_Bed

	if err := c.ShouldBindJSON(&mapbed); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&mapbed).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": mapbed})

}

// GET /user/:id

func GetMapBed(c *gin.Context) {
	var mapbed entity.Map_Bed
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM map_beds WHERE id = ?", id).Scan(&mapbed).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": mapbed})
}

// GET /users

func ListMapBed(c *gin.Context) {

	var mapbeds []entity.Map_Bed

	if err := entity.DB().Raw("SELECT * FROM map_beds").Scan(&mapbeds).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": mapbeds})

}

