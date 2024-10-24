package controllers

import (
	"ToDoList-rest-api/configs"
	"ToDoList-rest-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateAuthor(c *gin.Context) {

	var author models.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := configs.DB.Create(&author).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, author)
}

func GetAuthors(c *gin.Context) {

	var authors []models.Author
	configs.DB.Find(&authors)

	c.JSON(http.StatusOK, authors)
}

func GetAuthor(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	var author models.Author
	if err := configs.DB.Preload("Books").First(&author, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Author not found",
		})
		return
	}

	c.JSON(http.StatusOK, author)
}

func UpdateAuthor(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	var author models.Author
	if err := configs.DB.First(&author, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Author not found",
		})
		return
	}

	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	configs.DB.Save(&author)
	c.JSON(http.StatusOK, author)
}

func DeleteAuthor(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	if err := configs.DB.Delete(&models.Author{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Author deleted"})
}
