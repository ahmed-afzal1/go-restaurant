package main

import (
	_ "github.com/ahmed-afzal1/restaurant/cmd/restaurant/docs"
	"github.com/ahmed-afzal1/restaurant/config"
	"github.com/ahmed-afzal1/restaurant/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Swagger
//
//	@title						Restaurant Management API
//	@version					1.0
//	@description				A comprehensive API for managing restaurant data.
//	@contact.name				API Support Team
//	@license.name				MIT
//	@host						http://localhost:8080/
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
	r.Static("/assets", "./assets")

	r.Use(gin.Logger())
	config.ConnectDatabase()
	routers.RegisterRoutes(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run()
}
