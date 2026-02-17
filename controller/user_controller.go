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
	authService service.UserAuthService
}

func NewUserController(userService service.UserService, authService service.UserAuthService) *userController {
	return &userController{userService, authService}
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

	response := util.APIResponse(http.StatusOK, util.MessageSuccess, user.FormatUserRegisterResponse(newUser))
	c.JSON(http.StatusOK, response)
}

func (h *userController) LoginUser(c *gin.Context) {
	var inputUser user.UserLoginRequest

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

	loginUser, err := h.userService.LoginUser(&inputUser)
	if err != nil {

		response := util.APIError(
			http.StatusUnprocessableEntity,
			util.MessageFailedRegister,
			err.Error(),
		)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := h.authService.GenerateToken(loginUser.UserId, loginUser.Role)
	if err != nil {

		response := util.APIError(
			http.StatusUnprocessableEntity,
			util.MessageFailedRegister,
			err.Error(),
		)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := util.APIResponse(http.StatusOK, util.MessageSuccess, user.FormatUserLoginResponse(loginUser, token))
	c.JSON(http.StatusOK, response)
}
