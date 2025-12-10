package route

import (
	"go-gin-api/internal/controller"
	"go-gin-api/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	userService := service.NewUserService()
	userController := controller.NewUserController(userService)

	users := r.Group("/users")
	{
		users.GET("/", userController.GetUsers)
		users.POST("/register", userController.Register)
		users.GET("/:id", userController.GetUserByID)
		users.GET("/email/:email", userController.GetUserByEmail)
		users.GET("/phone/:phone", userController.GetUserByPhone)
	}
}
