package controller

import (
	"SaveMate/models/user"
	"SaveMate/util"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MockUserService struct {
	User *user.User
	Err  error
}

type MockUserAuthService struct {
	Token *jwt.Token
	Err   error
}

func (m *MockUserAuthService) GenerateToken(userID string, role string) (*jwt.Token, error) {
	return m.Token, m.Err
}

func (m *MockUserAuthService) ValidationToken(token string) (*jwt.Token, error) {
	return m.Token, m.Err
}

func (m *MockUserService) RegisterUser(input *user.UserRegister) (*user.User, error) {
	if m.Err != nil {
		return nil, m.Err
	}

	if m.User != nil {
		return nil, errors.New(util.MessageEmailIsNotAvailable)
	}

	return &user.User{
		Email:     input.Email,
		Username:  input.Username,
		UserId:    "123456",
		Role:      "USER",
		CreatedAt: time.Now(),
	}, nil
}

func (m *MockUserService) LoginUser(input *user.UserLoginRequest) (*user.User, error) {
	return nil, nil
}

func (m *MockUserService) FindByUserId(userID string) (*user.User, error) {
	return m.User, m.Err
}

type RegisterResponse struct {
	Status  int                       `json:"status"`
	Message string                    `json:"message"`
	Data    user.UserRegisterResponse `json:"data"`
}

func TestRegisterUserWhenEmailIsAvailable_expectedReturnCode422(t *testing.T) {
	router := gin.Default()

	mockService := &MockUserService{
		User: &user.User{
			Email: "test@gmail.com",
		},
		Err: nil,
	}

	controller := NewUserController(mockService, nil)
	router.POST("/register", controller.RegisterUser)

	body := `{"username" : "test", "email" : "test@gmail.com", "password" : "inipassword"}`
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, 422, w.Code)

}

func TestRegisterUserWhenEmailIsNotMatchWithFormat_expectedReturnCode422(t *testing.T) {
	router := gin.Default()

	mockService := &MockUserService{
		User: nil,
		Err:  nil,
	}

	controller := NewUserController(mockService, nil)
	router.POST("/register", controller.RegisterUser)

	body := `{"username" : "test", "email" : "test&gmail.com", "password" : "P@ssw0rd"}`
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	var resp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &resp)

	assert.Equal(t, 422, w.Code)
	assert.NotNil(t, resp["errors"])

}

func TestRegisterUser_expectedSuccess(t *testing.T) {
	router := gin.Default()

	mockService := &MockUserService{
		User: nil,
		Err:  nil,
	}

	controller := NewUserController(mockService, nil)
	router.POST("/register", controller.RegisterUser)

	body := `{"username" : "test", "email" : "test@gmail.com", "password" : "inipassword"}`
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

}

func TestRegisterUserWhenFieldRequiredNotFilled_expectedReturn422Code(t *testing.T) {
	router := gin.Default()

	mockService := &MockUserService{
		User: &user.User{
			Email:  "test@gmail.com",
			UserId: "123456",
			Role:   "USER",
		},
		Err: nil,
	}

	controller := NewUserController(mockService, nil)
	router.POST("/register", controller.RegisterUser)

	body := `{"email" : "test@gmail.com", "password" : "inipassword"}`
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, 422, w.Code)

}

func TestRegisterUserWhenSucces_expectedReturnRegisterResponseFormatter(t *testing.T) {
	router := gin.Default()

	mockService := &MockUserService{
		User: nil,
		Err:  nil,
	}

	controller := NewUserController(mockService, nil)
	router.POST("/register", controller.RegisterUser)

	body := `{"username" : "test", "email" : "test@gmail.com", "password" : "inipassword"}`
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	respBody, err := io.ReadAll(w.Body)

	assert.Nil(t, err)

	var resp RegisterResponse
	err = json.Unmarshal(respBody, &resp)
	assert.Nil(t, err)

	assert.Equal(t, 200, resp.Status)
	assert.Equal(t, util.MessageSuccess, resp.Message)
	assert.Equal(t, "test@gmail.com", resp.Data.Email)
	assert.Equal(t, "test", resp.Data.Username)
	assert.Equal(t, "123456", resp.Data.UserId)
	assert.Equal(t, "USER", resp.Data.Role)
}
