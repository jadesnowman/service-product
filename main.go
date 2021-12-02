package main

import (
	"fmt"
	"github.com/jadesnowman/service-product/controllers"
	"github.com/jadesnowman/service-product/db"
	"github.com/jadesnowman/service-product/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	router := gin.Default()
	v1 := router.Group("/api/v1")

	v1.POST("/auth/register", controllers.Register)
	v1.POST("/auth/login", controllers.Login)

	v1.GET("/users", controllers.GetUsers)

	v1.Use(middlewares.AuthHandler())
	{
		v1.GET("/products", controllers.Index)
		v1.POST("/products", controllers.Store)
		v1.GET("/products/:id", controllers.Show)
		v1.PUT("/products/:id", controllers.Update)
		v1.DELETE("/products/:id", controllers.Delete)
	}

	fmt.Println("serve at http://localhost:8082")

	router.Run(":8082")
}
