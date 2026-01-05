package main

import (
	"latihan-rest-api-gin-golang/database"
	"latihan-rest-api-gin-golang/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// db connection
	database.Database()

	//versioning
	v1 := router.Group("/v1")

	// path
	v1.GET("/", handler.GetHandler)
	v1.POST("/post", handler.PostHandler)
	v1.PATCH("/:id", handler.UpdateHandler)

	router.Run(":8000")
}