package db

import (
	"fmt"
	"service-product/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func InitDB() {
	dsn := "root:root@tcp(127.0.0.1:3306)/sandbox_service_product?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(&model.Product{})
}

func GetDB() *gorm.DB {
	fmt.Print("Nilai")
	return db
}