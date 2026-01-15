package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Database() *gorm.DB{
	dsn := "app:root@tcp(127.0.0.1:3306)/product?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("error connect to db")
	}
	fmt.Println("sucsess connect to db")
	DB = db	
	return db
}