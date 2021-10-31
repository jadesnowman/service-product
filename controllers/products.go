package controllers

import (
	"service-product/db"
	"service-product/model"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	database := db.GetDB()

	product := []model.Product{}

	result := database.Find(&product)

	if result.Error != nil {
		c.AbortWithStatus(404)
	} else {
		c.JSON(200, product)
	}
}

func Store(c *gin.Context) {
	database := db.GetDB()

	product := model.Product{
		Name:  "Macbook M2 2021",
		Code:  "PRO22M2",
		Price: 14990000,
	}

	result := database.Create(&product)

	if result.Error != nil {
		c.AbortWithStatus(404)
	} else {
		c.JSON(200, product)
	}
}
