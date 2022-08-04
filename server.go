package main

import (
	"github.com/WantToBePro31/bookstore-api/models"
	"github.com/WantToBePro31/bookstore-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()
	db := models.CreateDB()
	routes.SetRoutes(db, server)
	server.Run("localhost:8080")
}
