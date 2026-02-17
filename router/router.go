package router

import (
	"SaveMate/controller"
	"SaveMate/repository"
	"SaveMate/service"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {

	router := gin.Default()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	authService := service.NewUserAuthService()
	userController := controller.NewUserController(userService, authService)

	user := router.Group("/api/user")
	{
		user.POST("/register", userController.RegisterUser)
		user.POST("/login", userController.LoginUser)
	}

	return router
}
