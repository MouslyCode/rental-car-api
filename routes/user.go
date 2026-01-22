package routes

import (
	"github.com/MouslyCode/rental-car-api/controllers"
	"github.com/MouslyCode/rental-car-api/middlewares"
	"github.com/gin-gonic/gin"
)

func User(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.POST("/register", controllers.Register)
		users.POST("/login", controllers.Login)
		users.GET("/", middlewares.AuthMiddleware(), controllers.Get)
		users.PUT("/:id", middlewares.AuthMiddleware(), controllers.Update)
		users.DELETE("/:id", middlewares.AuthMiddleware(), controllers.Delete)
	}
}
