package routes

import (
	"ToDoList-rest-api/controllers"
	"github.com/gin-gonic/gin"
)

func SetRoutes(r *gin.Engine) {

	authorRoutes := r.Group("/authors")
	{
		authorRoutes.POST("/", controllers.CreateAuthor)
		authorRoutes.GET("/", controllers.GetAuthors)
		authorRoutes.GET("/:id", controllers.GetAuthor)
		authorRoutes.PUT("/:id", controllers.UpdateAuthor)
		authorRoutes.DELETE("/:id", controllers.DeleteAuthor)
	}

	bookRoutes := r.Group("/books")
	{
		bookRoutes.POST("/", controllers.CreateBook)
		bookRoutes.GET("/", controllers.GetBooks)
		bookRoutes.GET("/:id", controllers.GetBook)
		bookRoutes.PUT("/:id", controllers.UpdateBook)
		bookRoutes.DELETE("/:id", controllers.DeleteBook)
	}

}
