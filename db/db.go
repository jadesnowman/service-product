package db

import (
	"fmt"

	"github.com/jadesnowman/service-product/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func InitDB() {
	dsn := "root:saw@d1kab@tcp(backend_database_mysql)/sandbox_service_product?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(&model.Product{})
	db.AutoMigrate(&model.User{})
}

func GetDB() *gorm.DB {
	return db
}
