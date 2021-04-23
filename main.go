package main

import (
	"github.com/prasadadireddi/scytale-restapi/controllers"
	"github.com/prasadadireddi/scytale-restapi/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to database
	models.ConnectDatabase()

	// Routes
	r.GET("/workloads", controllers.GetWorkloads)
	r.POST("/workloads", controllers.CreateWorkload)
	r.PATCH("/workloads/:spiffeid", controllers.UpdateWorkload)
	r.DELETE("/workloads/:spiffeid", controllers.DeleteWorkload)

	// Run the server
	r.Run()
}
