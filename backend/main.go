package main

import (
	"github.com/watanprasai/sa-65-example/controller"
	"github.com/watanprasai/sa-65-example/entity"
	"github.com/gin-gonic/gin"
)

func main() {

	entity.SetupDatabase()
	r := gin.Default()
	r.Use(CORSMiddleware())

	// User Routes
	r.GET("/users", controller.ListUsers)
	r.GET("/user/:id", controller.GetUser)
	r.POST("/users", controller.CreateUser)
	r.PATCH("/users", controller.UpdateUser)
	r.DELETE("/users/:id", controller.DeleteUser)
	// Run the server

	// Create Symptom
	r.POST("/symptom", controller.CreateSymptom)
	//
	r.GET("/getSymptom", controller.ListSymptom)

	r.GET("/listMapbed", controller.ListMapBed)
	r.GET("/getMapbed", controller.GetMapBed)

	r.GET("/getLevel", controller.GetLevel)
	r.Run()

}

func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		  c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		  c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		  c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		  c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
  
		  if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		  }
		  c.Next()
	}
  }