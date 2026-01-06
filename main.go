package main

import (
	"fmt"
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
	v1.DELETE("/:id", handler.DeleteHandler)

	// router.Run(":8000")
	
	person := Person{Name: "rayhan"}	
	SayHello(person)
	animal := Animal{Name : "kucing"}
	SayHello(animal)
}

//interface == contract
type HasName interface{
	GetName()string
}

type Person struct{
	Name string
}

func (person Person) GetName() string{
	return person.Name
}

func SayHello(hasName HasName){
	fmt.Println("hellow", hasName.GetName())
}

// contoh2
type Animal struct{
	Name string
}

func (animal Animal) GetName() string{
	return animal.Name
}