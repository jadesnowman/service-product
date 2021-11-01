package controllers

import (
	"net/http"
	"service-product/db"
	"service-product/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	user := model.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, model.Fail{
			Message: err.Error(),
		})
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		c.JSON(http.StatusConflict, model.Fail{
			Message: err.Error(),
		})
		return
	}

	user.Password = string(hashPassword)

	database := db.GetDB()
	result := database.Create(&user)

	if err := result.Error; err != nil {
		c.JSON(http.StatusConflict, model.Fail{
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, user)
}

func Login(c *gin.Context) {
	user := model.User{}

	err := c.ShouldBindJSON(&user)
	if err != nil {
		fail := model.Fail{
			Message: "Your account does not exists!",
		}

		c.JSON(403, fail)
		return
	}

	c.JSON(200, user)
}
