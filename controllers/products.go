package controllers

import (
	"fmt"
	"net/http"
	"service-product/db"
	"service-product/model"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	product := []model.Product{}
	database := db.GetDB()
	result := database.Order("id desc").Find(&product)

	if result.Error != nil {
		c.AbortWithStatus(404)
	} else {
		c.JSON(200, product)
	}
}

func Store(c *gin.Context) {

	product := model.Product{}
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database := db.GetDB()
	result := database.Create(&product)

	if result.Error != nil {
		c.AbortWithStatus(404)
	} else {
		c.JSON(200, product)
	}
}

func Show(c *gin.Context) {
	id := c.Params.ByName("id")

	product := model.Product{}
	database := db.GetDB()
	result := database.First(&product, id)

	if err := result.Error; err != nil {
		responseFail := model.Fail{
			Message: err.Error(),
		}
		fmt.Println(result)
		c.JSON(404, responseFail)
	} else {
		c.JSON(200, product)
	}
}
