package handler

import (
	"latihan-rest-api-gin-golang/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RootHandler(c *gin.Context){
			c.JSON(http.StatusOK, gin.H{
			"message": "sucsess",
		})
}

func PostHandler(c *gin.Context){
	var products entity.Products
	
	// masukan input body ke struct
	err := c.ShouldBindBodyWithJSON(&products)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "sucsess post datas",
		"datas" : products,
	})
}