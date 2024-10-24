package controllers

import (
	"ToDoList-rest-api/configs"
	"ToDoList-rest-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateBook(c *gin.Context) {
	var input models.CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var author models.Author
	if err := configs.DB.Find(&author, input.AuthorID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Author not found",
		})
		return
	}

	book := models.Book{
		Title:    input.Title,
		ISBN:     input.ISBN,
		AuthorID: author.ID,
		Author:   author,
	}

	if err := configs.DB.Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	configs.DB.Preload("Author").Find(&book, book.ID)

	c.JSON(http.StatusOK, book)
}

func GetBooks(c *gin.Context) {
	var books []models.Book
	configs.DB.Preload("Author").Find(&books)
	c.JSON(http.StatusOK, books)
}

func GetBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var book models.Book
	if err := configs.DB.Preload("Author").First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Book not found",
		})
		return
	}
	c.JSON(http.StatusOK, book)
}

func UpdateBook(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	var book models.Book
	if err := configs.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Book not found",
		})
		return
	}

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	configs.DB.Save(&book)
	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	if err := configs.DB.Delete(&models.Book{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}
