package database

import (
	"fmt"
	"latihan-api-gorm/entity"
	"latihan-api-gorm/repository"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Database() {
	dsn := "app:root@tcp(127.0.0.1:3306)/product?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("error connect to db")
	}
	fmt.Println("sucsess connect to db")
	DB = db

	repositoryLayer := repository.NewRepository(db)
	product, err := repositoryLayer.FindAll()
	if err != nil{
		fmt.Println("cannot find all prod")
	}
	// for i := 0; i < len(product); i++{
	// 	fmt.Println("produk : ", product[i].Name)
	// }  
	for i, p := range product {
    fmt.Println("produk", i, p.Name)
}

	prod := entity.Product{
		Name: "kontolodon",
		Price: 20000,
	}
	prods, err := repositoryLayer.Create(prod)
	fmt.Println(prods)

	id := 1
	productById, err := repositoryLayer.FindByID(id)
	fmt.Println("product by id ", id, productById.Name )

}