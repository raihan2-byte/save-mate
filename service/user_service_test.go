package service

import (
	"SaveMate/models/user"
	"SaveMate/util"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

type MockUserRepository struct {
	User *user.User
	Err  error
}

func (m *MockUserRepository) FindByEmail(email string) (*user.User, error) {
	if m.User != nil && m.User.Email == email {
		return m.User, nil
	}

	return nil, nil
}

func (m *MockUserRepository) CreateUser(u *user.User) (*user.User, error) {
	return u, nil
}

func (m *MockUserRepository) FindByUserId(id string) (*user.User, error) {
	return m.User, nil
}

func TestUserService(t *testing.T) {

	t.Run("TestRegisterUser_expectedSuccess", func(t *testing.T) {
		mockRepository := &MockUserRepository{
			User: nil,
			Err:  nil,
		}

		service := NewUserService(mockRepository)
		input := &user.UserRegister{
			Email:    "test@gmail.com",
			Password: "P@ssw0rd",
		}

		result, err := service.RegisterUser(input)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "test@gmail.com", result.Email)
	})

	t.Run("TestRegisterUserIfPasswordDoesntHaveUppercase_expectedMessageFailed", func(t *testing.T) {
		mockRepository := &MockUserRepository{
			User: nil,
			Err:  nil,
		}

		service := NewUserService(mockRepository)
		input := &user.UserRegister{
			Email:    "test@gmail.com",
			Password: "inipassword",
		}

		result, err := service.RegisterUser(input)

		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.Equal(t, "password must contain at least one uppercase letter", err.Error())
	})

	t.Run("TestRegisterWhenEmailExist_expectedFail", func(t *testing.T) {
		mockRepository := &MockUserRepository{
			User: &user.User{
				Email: "test@gmail.com",
			},
			Err: nil,
		}

		service := NewUserService(mockRepository)
		input := &user.UserRegister{
			Email:    "test@gmail.com",
			Password: "123",
		}

		result, err := service.RegisterUser(input)

		assert.NotNil(t, err)
		assert.Nil(t, result)
	})

	t.Run("TestLoginUser_expectedSuccess", func(t *testing.T) {
		hashedPassword, err := bcrypt.GenerateFromPassword(
			[]byte("P@ssw0rd"),
			bcrypt.DefaultCost,
		)

		mockRepository := &MockUserRepository{
			User: &user.User{
				Email:    "test@gmail.com",
				Password: string(hashedPassword),
			},
			Err: nil,
		}

		service := NewUserService(mockRepository)
		input := &user.UserLoginRequest{
			Email:    "test@gmail.com",
			Password: "P@ssw0rd",
		}

		result, err := service.LoginUser(input)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "test@gmail.com", result.Email)
	})

	t.Run("TestLoginUserWhenEmailIsNotFound_expectedFailed", func(t *testing.T) {
		mockRepository := &MockUserRepository{
			User: nil,
			Err:  nil,
		}

		service := NewUserService(mockRepository)
		input := &user.UserLoginRequest{
			Email:    "test@gmail.com",
			Password: "P@ssw0rd",
		}

		result, err := service.LoginUser(input)

		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.Equal(t, util.MessageAuthenticationFailed, err.Error())
	})

	t.Run("TestLoginUserWhenPasswordIsNotMatch_expectedFailed", func(t *testing.T) {
		hashedPassword, err := bcrypt.GenerateFromPassword(
			[]byte("P@ssw0rd"),
			bcrypt.DefaultCost,
		)

		mockRepository := &MockUserRepository{
			User: &user.User{
				Email:    "test@gmail.com",
				Password: string(hashedPassword),
			},
			Err: nil,
		}

		service := NewUserService(mockRepository)
		input := &user.UserLoginRequest{
			Email:    "test@gmail.com",
			Password: "P@ssw0rd1234",
		}

		result, err := service.LoginUser(input)

		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.Equal(t, util.MessageAuthenticationFailed, err.Error())
	})
}
