package middleware

import (
	"SaveMate/models/user"
	"SaveMate/service"
	"SaveMate/util"
	"errors"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(authService service.UserAuthService, userService service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if !strings.HasPrefix(authHeader, "Bearer ") {
			response := util.APIResponse(http.StatusUnauthorized, util.MessageUnauthorized, errors.New(util.MessageAuthenticationFailed))
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := strings.Split(authHeader, " ")[1]

		token, err := authService.ValidationToken(tokenString)

		if err != nil || !token.Valid {
			response := util.APIResponse(http.StatusUnauthorized, util.MessageUnauthorized, errors.New(util.MessageAuthenticationFailed))
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			response := util.APIResponse(http.StatusUnauthorized, util.MessageUnauthorized, errors.New(util.MessageAuthenticationFailed))
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID, ok := claims["user_id"].(string)

		if !ok {
			response := util.APIResponse(http.StatusUnauthorized, util.MessageUnauthorized, errors.New(util.MessageAuthenticationFailed))
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		role, ok := claims["role"].(string)

		if !ok {
			response := util.APIResponse(http.StatusUnauthorized, util.MessageUnauthorized, errors.New(util.MessageAuthenticationFailed))
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		user, err := userService.FindByUserId(userID)

		if err != nil {
			response := util.APIResponse(http.StatusUnauthorized, util.MessageUnauthorized, errors.New(util.MessageAuthenticationFailed))
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
		c.Set("role", role)

		c.Next()
	}
}

func RoleMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {

		userRole := c.GetString("role")

		if userRole != role {

			response := util.APIResponse(http.StatusForbidden, util.MessageForbidden, errors.New(util.MessageForbidden))

			c.AbortWithStatusJSON(http.StatusForbidden, response)

			return
		}

		c.Next()
	}
}

func CurrentUser(c *gin.Context) *user.User {

	currentUser, exists := c.Get("currentUser")

	if !exists {
		return nil
	}

	return currentUser.(*user.User)

}
