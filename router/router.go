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
	userController := controller.NewUserController(userService)

	user := router.Group("/api/user")
	{
		user.POST("/register", userController.RegisterUser)
	}

	return router
}
