package routes

import (
	"github.com/BlazeSyn/bookstore-api/controllers"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetRoutes(db *gorm.DB, r *gin.Engine) {
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.GET("/books", controllers.GetAllBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.GetBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("books/:id", controllers.RemoveBook)
}
