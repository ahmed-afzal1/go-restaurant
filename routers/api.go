package routers

import (
	"github.com/ahmed-afzal1/restaurant/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	authRoute := r.Group("/v1/auth")
	{
		authRoute.POST("/register", controllers.Register)
		authRoute.POST("/login", controllers.Login)
	}

	r.GET("/dashboard", controllers.Dashboard)
}
