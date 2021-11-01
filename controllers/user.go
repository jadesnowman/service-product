package controllers

import (
	"service-product/db"
	"service-product/model"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	user := []model.User{}

	database := db.GetDB()
	database.Order("id desc").Find(&user)

	c.JSON(200, user)
}
