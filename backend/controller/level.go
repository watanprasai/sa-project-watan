package controller

import (
	"github.com/watanprasai/sa-65-example/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

// POST /users

func CreateLevel(c *gin.Context) {

	var level entity.Level

	if err := c.ShouldBindJSON(&level); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&level).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": level})

}

// GET /user/:id

func GetLevel(c *gin.Context) {

	var level entity.Level

	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM levels WHERE id = ?", id).Scan(&level).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": level})

}

// GET /users

func ListLevel(c *gin.Context) {

	var levels []entity.Level

	if err := entity.DB().Raw("SELECT * FROM levels").Scan(&levels).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": levels})

}
