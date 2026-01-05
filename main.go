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
	v1.GET("/", handler.RootHandler)
	v1.POST("/post", handler.PostHandler)





	router.Run(":8000")
}