package routes

import (
	"github.com/HrushikeshAnandSarangi/go-rest/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Register)
	r.GET("", controllers.Register)
}
