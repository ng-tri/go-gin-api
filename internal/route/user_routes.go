package route

import (
	"go-gin-api/internal/handler"
	"go-gin-api/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	userService := service.NewUserService()
	userHandler := handler.NewUserHandler(userService)

	users := r.Group("/users")
	{
		users.GET("/", userHandler.GetUsers)
		users.POST("/register", userHandler.Register)
		users.GET("/:id", userHandler.GetUserByID)
		users.GET("/email/:email", userHandler.GetUserByEmail)
		users.GET("/phone/:phone", userHandler.GetUserByPhone)
	}
}
