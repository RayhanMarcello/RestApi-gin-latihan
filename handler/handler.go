package handler

import (
	"latihan-api-gorm/database"
	"latihan-api-gorm/entity"

	"github.com/gin-gonic/gin"
)

func HandleGetProduct(c *gin.Context){
	var product []entity.Product

	err := database.DB.Find(&product).Error
	if err != nil {
		c.JSON(500, gin.H{
			"err" : err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"message" : "OK",
		"datas" : product,
	})
}

func HandlerPostProduct(c *gin.Context){
	var product entity.Product
	
	err := c.ShouldBindJSON(&product)
	if err != nil {
		c.JSON(500, gin.H{
			"message" : "server error",
		})
	}

	if err := database.DB.Create(&product).Error; err != nil {
		c.JSON(500, gin.H{
			"message" : err.Error(),
		})
	}

	if err := database.DB.Last(&product).Error; err != nil {
		c.JSON(500, gin.H{
			"message" : err.Error(),
		})
		return
	}
	c.JSON(500, gin.H{
			"datas" : product, 
	})

}