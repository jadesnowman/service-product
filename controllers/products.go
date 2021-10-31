package controllers

import (
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
		c.JSON(404, responseFail)
	} else {
		c.JSON(200, product)
	}
}

func Update(c *gin.Context) {
	id := c.Params.ByName("id")

	product := model.Product{}
	database := db.GetDB()
	result := database.Where("id = ?", id).First(&product)

	if err := result.Error; err != nil {
		responseFail := model.Fail{
			Message: err.Error(),
		}

		c.JSON(404, responseFail)
	} else {
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		database.Save(&product)

		c.JSON(200, product)
	}
}

func Delete(c *gin.Context) {
	id := c.Params.ByName("id")

	product := model.Product{}
	database := db.GetDB()
	result := database.Delete(&product, id)

	if err := result.Error; err != nil {
		responseFail := model.Fail{
			Message: err.Error(),
		}

		c.JSON(404, responseFail)
	} else {
		c.JSON(200, model.Success{
			Message: "Data successfully deleted",
		})
	}
}
