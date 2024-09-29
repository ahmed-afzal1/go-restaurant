package main

import (
	"github.com/ahmed-afzal1/restaurant/config"
	"github.com/ahmed-afzal1/restaurant/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	r := gin.Default()

	r.Use(gin.Logger())
	config.ConnectDatabase()
	routers.RegisterRoutes(r)

	r.Run()
}
