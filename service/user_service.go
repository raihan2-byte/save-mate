package service

import (
	"SaveMate/models/user"
	"SaveMate/repository"
	"SaveMate/util"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(user *user.UserRegister) (*user.User, error)
	// FindByUserId(UserId int) (*user.User, error)
}

type userService struct {
	repositoryUser repository.UserRepository
}

func NewUserService(repositoryUser repository.UserRepository) *userService {
	return &userService{repositoryUser}
}

func (s *userService) RegisterUser(userRequest *user.UserRegister) (*user.User, error) {

	cek, err := s.repositoryUser.FindByEmail(userRequest.Email)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if cek != nil {
		return nil, errors.New(util.MessageEmailIsAvailable)
	}

	newUser := &user.User{}
	newUser.UserId = uuid.New().String()
	newUser.Username = userRequest.Username
	newUser.Email = userRequest.Email
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.MinCost)
	if err != nil {
		return newUser, err
	}
	newUser.Password = string(passwordHash)
	newUser.Role = util.RoleUser
	newUser.CreatedAt = time.Now()
	newUser.UpdatedAt = time.Now()

	registerUser, err := s.repositoryUser.CreateUser(newUser)
	if err != nil {
		return registerUser, err
	}
	return registerUser, nil
}
