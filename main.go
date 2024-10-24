package main

import (
	"ToDoList-rest-api/configs"
	"ToDoList-rest-api/models"
	"ToDoList-rest-api/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	initDb()

	r := gin.Default()
	routes.SetRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func initDb() {
	dsn := "host=localhost user=postgres password=password dbname=bookstore_api port=5432 sslmode=disable TimeZone=Africa/Nairobi"

	var err error
	configs.DB, err = configs.ConnectDb(dsn)
	if err != nil {
		log.Fatal("Connection to db Failed", err)
	}

	if err = configs.DB.AutoMigrate(&models.Author{}, &models.Book{}); err != nil {
		log.Fatal("Migration Failed", err)
	}
}
