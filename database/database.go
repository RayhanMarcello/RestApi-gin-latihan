package database

import (
	"fmt"

	"latihan-rest-api-gin-golang/handler"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
func Database(){
	
	dsn := "app:root@tcp(127.0.0.1:3306)/latihanConnectDb?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("error connect db")
	}
	DB = db
	fmt.Println("sucsess connect db")
	db.AutoMigrate(&handler.Model{})
}