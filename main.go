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
	r.POST("/products", controllers.Store)
	r.GET("/products/:id", controllers.Show)
	r.PUT("/products/:id", controllers.Update)
	r.DELETE("/products/:id", controllers.Delete)

	r.Run(":8082")
}
