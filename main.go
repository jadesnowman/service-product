package main

import (
	"service-product/controllers"
	"service-product/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	r := gin.Default()

	r.GET("/products", controllers.Index)
	r.GET("/products/store", controllers.Store)

	r.Run(":8082")
}
