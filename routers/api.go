package routers

import (
	"github.com/ahmed-afzal1/restaurant/controllers"
	"github.com/ahmed-afzal1/restaurant/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	authRoute := r.Group("/api/v1/auth")
	{
		authRoute.POST("/register", controllers.Register)
		authRoute.POST("/login", controllers.Login)
	}

	userRoute := r.Group("/api/v1/admin")
	userRoute.Use(middleware.AdminAuthenticate())
	{
		userRoute.GET("/details", controllers.Details)
	}

	categoryRoute := r.Group("/api/v1/admin/category")
	categoryRoute.Use(middleware.AdminAuthenticate())
	{
		categoryRoute.GET("/index", controllers.GetAllCategories)
		categoryRoute.POST("/store", controllers.CategoryStore)
		categoryRoute.GET("/edit/:id", controllers.CategoryEdit)
		categoryRoute.GET("/status/:id/:status", controllers.CategoryStatusUpdate)
		categoryRoute.PATCH("/update/:id", controllers.CategoryUpdate)
		categoryRoute.DELETE("delete/:id", controllers.CategoryDelete)
	}
}
