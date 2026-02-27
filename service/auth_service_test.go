package service

import (
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func TestAuthService(t *testing.T) {
	t.Run("TestGenerateToken_ExpectedSuccess", func(t *testing.T) {
		authService := NewUserAuthService()
		userID := "user123"
		role := "admin"

		tokenString, err := authService.GenerateToken(userID, role)

		assert.Nil(t, err)
		assert.NotEmpty(t, tokenString)

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return SECRET_KEY, nil
		})

		assert.Nil(t, err)
		assert.True(t, token.Valid)

		claims := token.Claims.(jwt.MapClaims)
		assert.Equal(t, "user123", claims["user_id"])
		assert.Equal(t, "admin", claims["role"])
	})

	t.Run("TestGenerateTokenWhenUserIdNil_ExpectedFailed", func(t *testing.T) {
		authService := NewUserAuthService()
		userID := ""
		role := "admin"

		tokenString, err := authService.GenerateToken(userID, role)

		assert.Equal(t, "", tokenString)
		assert.NotNil(t, err)

	})

	t.Run("TestGenerateTokenWhenRoleIsNil_ExpectedFailed", func(t *testing.T) {
		authService := NewUserAuthService()
		userID := "123456rthrhd"
		role := ""

		tokenString, err := authService.GenerateToken(userID, role)

		assert.Equal(t, "", tokenString)
		assert.NotNil(t, err)

	})

	t.Run("TestValidationToken_ExpectedSuccess", func(t *testing.T) {
		authService := NewUserAuthService()
		userID := "user123"
		role := "admin"

		tokenString, err := authService.GenerateToken(userID, role)

		assert.Nil(t, err)
		assert.NotNil(t, tokenString)

		token, err := authService.ValidationToken(tokenString)

		assert.Nil(t, err)
		assert.NotNil(t, token)
		assert.True(t, token.Valid)

		claims := token.Claims.(jwt.MapClaims)

		assert.Equal(t, "user123", claims["user_id"])
		assert.Equal(t, "admin", claims["role"])
	})

	t.Run("TestValidationTokenInvalidSecretKey_ExpectedFailed", func(t *testing.T) {
		authService := NewUserAuthService()
		userID := "user123"
		role := "admin"

		tokenString, err := authService.GenerateToken(userID, role)

		assert.Nil(t, err)
		assert.NotNil(t, tokenString)

		SECRET_KEY = []byte("secret-salah")
		token, err := authService.ValidationToken(tokenString)

		assert.Nil(t, token)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "signature")
	})
}
