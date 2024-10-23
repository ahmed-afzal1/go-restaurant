package main

import (
	"github.com/ahmed-afzal1/restaurant/config"
	_ "github.com/ahmed-afzal1/restaurant/docs"
	"github.com/ahmed-afzal1/restaurant/routers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"time"
)

// Swagger
//
//	@title						Restaurant Management API
//	@version					1.0
//	@description				A comprehensive API for managing restaurant data.
//	@contact.name				API Support Team
//	@license.name				MIT
//	@host						localhost:8081
//	@BasePath					/api/v1
//	@schemes					http https
//	@securityDefinitions.apiKey	JWT
//	@in							header
//	@name						Authorization
//	@description				JWT security accessToken. Please add it in the format "Bearer {AccessToken}" to authorize your requests.
func main() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	r := gin.Default()

	// CORS middleware setup
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080", "https://localhost:8080"}, // Add your domain or localhost
		AllowMethods:     []string{"POST", "GET", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Static("/assets", "./assets")

	r.Use(gin.Logger())
	config.ConnectDatabase()
	routers.RegisterRoutes(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run(":8081")
}
