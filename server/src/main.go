package main

import (
	"github.com/gin-gonic/gin"
	"h-vistars/server/src/controllers"
	"h-vistars/server/src/database"
)

func main() {
	r := gin.Default()

	database.ConnectDatabase()

	var api = r.Group("/api")

	api.GET("/beats", controllers.GetBeats)
	api.GET("/beats/:id", controllers.GetBeatByID)
	api.POST("/beats", controllers.CreateBeat)
	api.PATCH("/beats/:id", controllers.UpdateBeat)
	api.DELETE("/beats/:id", controllers.DeleteBeat)

	err := r.Run()
	if err != nil {
		return
	}
}
