package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/ahmed-afzal1/restaurant/requests"
	"github.com/ahmed-afzal1/restaurant/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Dashboard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "I am from dashboard"})
}

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

	data, _ := services.Signup(req)

	c.JSON(http.StatusAccepted, gin.H{"data": data})
}

func Login(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{"message": "Login successfully"})
}
