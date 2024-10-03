package routers

import (
	"github.com/ahmed-afzal1/restaurant/controllers"
	"github.com/ahmed-afzal1/restaurant/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	authRoute := r.Group("/v1/auth")
	{
		authRoute.POST("/register", controllers.Register)
		authRoute.POST("/login", controllers.Login)
	}

	userRoute := r.Group("/v1/admin")

	userRoute.Use(middleware.AdminAuthenticate())
	{
		userRoute.GET("/details", controllers.Details)
	}

	categoryRoute := r.Group("/v1/admin/category")

	categoryRoute.Use(middleware.AdminAuthenticate())
	{
		categoryRoute.GET("/index", controllers.GetAllCategories)
	}
}
