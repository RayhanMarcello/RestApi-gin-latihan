package handler

import (
	"fmt"
	"latihan-rest-api-gin-golang/database"
	"latihan-rest-api-gin-golang/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHandler(c *gin.Context){
	var model []entity.Model
	err := database.DB.Find(&model)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
		"message" : "sucsess",
		"serverMessage" : err,
	})
}
	c.JSON(http.StatusOK, gin.H{
		"datas" : model,
	})

}

// func PostHandler(c *gin.Context){
// 	var products entity.Products
	
// 	// masukan input body ke struct
// 	err := c.ShouldBindBodyWithJSON(&products)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message" : err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"message" : "sucsess post datas",
// 		"datas" : products,
// 	})
// }

// db

func PostHandler(c *gin.Context){
	var model entity.Model
	
	models := entity.Model{
		Name: model.Name,
	}

	// masukan input body ke struct
	err := c.ShouldBindJSON(&models)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : err.Error(),
		})
		return
	}

	// db
	if err := database.DB.Create(&models).Error; err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message" : err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "sucsess post datas",
		"datas" : models,
	})
}

func UpdateHandler(c *gin.Context){
	var model entity.Model

	var id = c.Param("id")
	// cek ke db datanya ada apa ga
	err := database.DB.Find(&model, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : "didnt fount data with id : %v, id" ,
			"mesServer" : err.Error(),
		})
		return
	}
	
	// reweite ke struct dari body req
	if err := c.ShouldBindJSON(&model);err != nil {
		fmt.Println(err.Error())
		return
	}

	// simpan ke db gw
	if err := database.DB.Save(&model).Error;err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"mesServer" : err,
		})
		return
	}
	
	// tampilkan update
	c.JSON(http.StatusOK, gin.H{
		"message" : "sucsess update data by id",
		"datas" : model,
	})


}


