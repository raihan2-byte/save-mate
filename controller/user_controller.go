package controller

import (
	"SaveMate/models/user"
	"SaveMate/service"
	"SaveMate/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService service.UserService
	// authService auth.UserAuthService
}

func NewUserController(userService service.UserService /*, authService auth.UserAuthService*/) *userController {
	return &userController{userService /*, authService*/}
}

func (h *userController) RegisterUser(c *gin.Context) {
	var inputUser user.UserRegister

	err := c.ShouldBindJSON(&inputUser)
	if err != nil {

		response := util.APIError(
			http.StatusUnprocessableEntity,
			util.MessageFailedRegister,
			err.Error(),
		)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(&inputUser)
	if err != nil {

		response := util.APIError(
			http.StatusUnprocessableEntity,
			util.MessageFailedRegister,
			err.Error(),
		)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := util.APIResponse(http.StatusOK, util.MessageSuccessRegister, user.FormatUserRegisterResponse(newUser))
	c.JSON(http.StatusOK, response)
}
