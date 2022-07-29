package controllers

import (
	"net/http"

	"github.com/BlazeSyn/bookstore-api/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type BookInput struct {
	Title        string `json:"title"`
	Author       string `json:"author"`
	Release_Year int    `json:"release_year"`
}

func GetAllBooks(c *gin.Context) {
	books := new([]models.Book)
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Find(&books).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": books})
}

func CreateBook(c *gin.Context) {
	req := new(BookInput)
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	book := models.Book{
		Title:        req.Title,
		Author:       req.Author,
		Release_Year: req.Release_Year,
	}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func GetBook(c *gin.Context) {
	book := new(models.Book)
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func UpdateBook(c *gin.Context) {
	book := new(models.Book)
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	req := new(BookInput)
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	db.Model(&book).Update(req)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func RemoveBook(c *gin.Context) {
	book := new(models.Book)
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	db.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"data": book})
}
