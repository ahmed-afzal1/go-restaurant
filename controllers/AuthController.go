package controllers

import (
	"context"
	"github.com/ahmed-afzal1/restaurant/config"
	"github.com/ahmed-afzal1/restaurant/models"
	"github.com/ahmed-afzal1/restaurant/requests"
	"github.com/ahmed-afzal1/restaurant/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

func Register(c *gin.Context) {
	var req requests.RegisterRequest
	var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	if err := c.ShouldBind(&req); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			formattedErrors := requests.FormatValidationError(validationErrors)
			c.JSON(http.StatusBadRequest, gin.H{"errors": formattedErrors})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	savedUser, err := services.Signup(req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": savedUser})
}

func Login(c *gin.Context) {
	var req requests.LoginRequest

	var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	if err := c.ShouldBind(&req); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			formattedErrors := requests.FormatValidationError(validationErrors)
			c.JSON(http.StatusBadRequest, gin.H{"errors": formattedErrors})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	loginUser, err := services.Login(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "", "error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"data": loginUser, "error": ""})
}

func Details(c *gin.Context) {
	email, exists := c.Get("email")
	var user models.User
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	emailStr, ok := email.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email format"})
		return
	}

	getUser := config.DB.Where("email =?", emailStr).First(&user)

	if getUser.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": getUser.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": user})
}
