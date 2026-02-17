package service

import (
	"SaveMate/models/user"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestRegisterUser_expectedSuccess(t *testing.T) {
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
}

func TestRegisterUserIfPasswordDoesntHaveUppercase_expectedMessageFailed(t *testing.T) {
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

}

func TestRegisterWhenEmailExist_expectedFail(t *testing.T) {
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
}
