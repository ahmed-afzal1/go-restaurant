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

	adminRoute := r.Group("/api/v1/admin")
	adminRoute.Use(middleware.AdminAuthenticate())
	{
		adminRoute.GET("/details", controllers.Details)

		categoryRoute := adminRoute.Group("/category")
		{
			categoryRoute.GET("/index", controllers.GetAllCategories)
			categoryRoute.POST("/store", controllers.CategoryStore)
			categoryRoute.GET("/edit/:id", controllers.CategoryEdit)
			categoryRoute.GET("/status/:id/:status", controllers.CategoryStatusUpdate)
			categoryRoute.PATCH("/update/:id", controllers.CategoryUpdate)
			categoryRoute.DELETE("delete/:id", controllers.CategoryDelete)
		}

		subCategoryRoute := adminRoute.Group("/sub-category")
		{
			subCategoryRoute.GET("/index", controllers.GetAllSubCategories)
			// subCategoryRoute.POST("/store", controllers.CategoryStore)
			// subCategoryRoute.GET("/edit/:id", controllers.CategoryEdit)
			// subCategoryRoute.GET("/status/:id/:status", controllers.CategoryStatusUpdate)
			// subCategoryRoute.PATCH("/update/:id", controllers.CategoryUpdate)
			// subCategoryRoute.DELETE("delete/:id", controllers.CategoryDelete)
		}
	}
}
