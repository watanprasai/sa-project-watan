package controller

import (
	"github.com/watanprasai/sa-65-example/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {

	var user entity.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})

}


func GetUser(c *gin.Context) {

	var user entity.User

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM users WHERE id = ?", id).Scan(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})

}


func ListUsers(c *gin.Context) {

	var users []entity.User

	if err := entity.DB().Raw("SELECT * FROM users").Scan(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})

}


func DeleteUser(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM users WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})

}


func UpdateUser(c *gin.Context) {

	var user entity.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", user.ID).First(&user); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	if err := entity.DB().Save(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}


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


func GetMapBed(c *gin.Context) {
	var mapbed entity.Map_Bed
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM map_beds WHERE id = ?", id).Scan(&mapbed).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": mapbed})
}


func ListMapBed(c *gin.Context) {

	var mapbeds []entity.Map_Bed

	if err := entity.DB().Raw("SELECT * FROM map_beds").Scan(&mapbeds).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": mapbeds})

}


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


func GetLevel(c *gin.Context) {

	var level entity.Level

	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM levels WHERE id = ?", id).Scan(&level).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": level})

}


func ListLevel(c *gin.Context) {

	var levels []entity.Level

	if err := entity.DB().Raw("SELECT * FROM levels").Scan(&levels).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": levels})

}



func CreateSymptom(c *gin.Context) {

	var user entity.User
	var level entity.Level
	var mapbed entity.Map_Bed
	var symptom entity.Symptom

	// ?????????????????????????????????????????????????????????????????????????????? 8 ??????????????? bind ?????????????????????????????? symptom
	if err := c.ShouldBindJSON(&symptom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ??????????????? map_bed ???????????? id
	if tx := entity.DB().Where("id = ?", symptom.MapbID).First(&mapbed); tx.RowsAffected == 0 {
    	c.JSON(http.StatusBadRequest, gin.H{"error": "map bed not found"})
    	return
    }

	// 10: ??????????????? user ???????????? id
	if tx := entity.DB().Where("id = ?", symptom.CheckID).First(&user); tx.RowsAffected == 0 {
    	c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
    	return
    }

	// 11: ??????????????? level ???????????? id
	if tx := entity.DB().Where("id = ?", symptom.LevelID).First(&level); tx.RowsAffected == 0 {
    	c.JSON(http.StatusBadRequest, gin.H{"error": "level not found"})
    	return
    }

	// 12: ??????????????? Symptom
    s := entity.Symptom{
		Check_date: 		symptom.Check_date,    // ???????????????????????????????????? check_date
		Temperature: 		symptom.Temperature,
		Pressure:			symptom.Pressure,
		Heart_rate:			symptom.Heart_rate,
		Comment:			symptom.Comment,
		Medicine:			symptom.Medicine,
    	Check:   		 	user,            // ?????????????????????????????????????????????????????? Entity User
    	Level:              level,               // ?????????????????????????????????????????????????????? Entity Level
    	Mapb:               mapbed,           // ?????????????????????????????????????????????????????? Entity Map_Bed
    }
	
	// 13: ??????????????????
    if err := entity.DB().Create(&s).Error; err != nil {
    	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    	return
    }
    c.JSON(http.StatusOK, gin.H{"data": s})

}



func GetSymptom(c *gin.Context) {

	var symptom entity.Symptom
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM users WHERE id = ?", id).Scan(&symptom).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": symptom})

}



func ListSymptom(c *gin.Context) {

	var symptoms []entity.Symptom

	if err := entity.DB().Raw("SELECT * FROM symptoms").Scan(&symptoms).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": symptoms})

}




