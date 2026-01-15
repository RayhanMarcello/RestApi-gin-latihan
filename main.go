package main

import (
	"latihan-api-gorm/database"
	"latihan-api-gorm/service"

	// "latihan-api-gorm/entity"
	"latihan-api-gorm/handler"
	"latihan-api-gorm/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	
	db := database.Database()

	repositoryLayer := repository.NewRepository(db)
	serviceLayer := service.NewProductService(repositoryLayer)
	h := handler.NewProductHandler(serviceLayer)

	r.GET("/product", h.GetAll)
	r.GET("/product/:id", h.GetAllById)
	r.POST("/product", h.Create)
	r.DELETE("/product/:id", h.DeleteById)

	r.Run(":8080")
}