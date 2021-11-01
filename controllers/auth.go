package controllers

import (
	"net/http"
	"service-product/db"
	"service-product/middlewares"
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

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, model.Fail{
			Message: err.Error(),
		})
		return
	}

	originalPassword := user.Password

	database := db.GetDB()
	result := database.Where("Email = ?", user.Email).First(&user)
	if err := result.Error; err != nil {
		c.JSON(http.StatusConflict, model.Fail{
			Message: err.Error(),
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(originalPassword))
	if err != nil {
		c.JSON(http.StatusConflict, model.Fail{
			Message: err.Error(),
		})
		return
	}

	token, err := middlewares.GenerateToken(int(user.ID), user.Email)
	if err != nil {
		c.JSON(http.StatusConflict, model.Fail{
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"token": token, "expires_at": middlewares.JWT_EXPIRES_AT})
}
