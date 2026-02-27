package middleware

import (
	"SaveMate/models/user"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MockAuthService struct {
	Token *jwt.Token
	Err   error
}

func (m *MockAuthService) ValidationToken(token string) (*jwt.Token, error) {
	return m.Token, m.Err
}

func (m *MockAuthService) GenerateToken(userID string, role string) (string, error) {
	return "token", m.Err
}

type MockUserService struct {
	User *user.User
	Err  error
}

func (m *MockUserService) FindByUserId(id string) (*user.User, error) {
	return m.User, m.Err
}
func (m *MockUserService) RegisterUser(userRequest *user.UserRegister) (*user.User, error) {
	return m.User, m.Err
}
func (m *MockUserService) LoginUser(loginRequest *user.UserLoginRequest) (*user.User, error) {
	return m.User, m.Err
}

func TestAuthMiddleware(t *testing.T) {
	t.Run("TestAuthMiddleware_Success", func(t *testing.T) {
		gin.SetMode(gin.TestMode)

		claims := jwt.MapClaims{
			"user_id": "123",
			"role":    "USER",
		}

		token := &jwt.Token{
			Valid:  true,
			Claims: claims,
		}

		mockAuth := &MockAuthService{
			Token: token,
			Err:   nil,
		}

		mockUserService := &MockUserService{
			User: &user.User{
				UserId: "123",
			},
			Err: nil,
		}

		router := gin.New()
		router.Use(AuthMiddleware(mockAuth, mockUserService))

		router.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "success"})
		})

		req, _ := http.NewRequest("GET", "/test", nil)
		req.Header.Set("Authorization", "Bearer validtoken")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("TestAuthMiddleware_NoHeader", func(t *testing.T) {
		gin.SetMode(gin.TestMode)

		mockAuth := &MockAuthService{}
		mockUserService := &MockUserService{}

		router := gin.New()
		router.Use(AuthMiddleware(mockAuth, mockUserService))

		router.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "success"})
		})

		req, _ := http.NewRequest("GET", "/test", nil)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
}
