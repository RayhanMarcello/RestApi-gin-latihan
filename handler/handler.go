package handler

import (
	"latihan-api-gorm/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct{
	service service.ProductService
}

func NewProductHandler(p service.ProductService) *ProductHandler{
	return &ProductHandler{p}
} 

func (h *ProductHandler) GetAll(c *gin.Context){
	 prod, err := h.service.GetAll()
	 if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err,
		})
		return
	 }

	 c.JSON(http.StatusOK, gin.H{
		"message" : "sucsess", 
		"datas" : prod,
	 })
}

type reqBody struct{
	Name string
	Price int 
}

func (h *ProductHandler) Create(c *gin.Context){
	var req reqBody
	// terima req body dari client
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : err,
		})
		return
	}
	//
	create, err := h.service.Create(req.Name, req.Price)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err" : err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message" : "sucsess create", 
		"data": create,
	})
}

func (h *ProductHandler) GetAllById(c *gin.Context){
	idStr := c.Param("id")
	id,err := strconv.ParseInt(idStr,10,64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err" : err,
		})
		return
	}

	product, err := h.service.GetAllById(int(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err" : err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message" : "sucsess get data by id",
		"data" : product,
	})
}
