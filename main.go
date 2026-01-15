package main

import (
	"latihan-api-gorm/database"
	"latihan-api-gorm/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database.Database()

	r.GET("/product", handler.HandleGetProduct)
	r.POST("/product", handler.HandlerPostProduct)

	r.Run(":8080")
}