package controllers

import (
	"fmt"
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

func Show(c *gin.Context) {
	id := c.Params.ByName("id")

	database := db.GetDB()

	product := model.Product{}

	result := database.First(&product, id)

	if result.Error != nil {
		responseFail := model.Fail{
			Message: result.Error.Error(),
		}
		fmt.Println(result)
		c.JSON(404, responseFail)
	} else {
		c.JSON(200, product)
	}
}
