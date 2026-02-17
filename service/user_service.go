package service

import (
	"SaveMate/models/user"
	"SaveMate/repository"
	"SaveMate/util"
	"database/sql"
	"errors"
	"time"
	"unicode"

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

func ValidatePassword(password string) error {
	var (
		hasUpper   bool
		hasLower   bool
		hasNumber  bool
		hasSpecial bool
	)

	if len(password) < 6 {
		return errors.New("password must be at least 6 characters")
	}

	for _, ch := range password {
		switch {
		case unicode.IsUpper(ch):
			hasUpper = true
		case unicode.IsLower(ch):
			hasLower = true
		case unicode.IsNumber(ch):
			hasNumber = true
		case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
			hasSpecial = true
		}
	}

	if !hasUpper {
		return errors.New("password must contain at least one uppercase letter")
	}
	if !hasLower {
		return errors.New("password must contain at least one lowercase letter")
	}
	if !hasNumber {
		return errors.New("password must contain at least one number")
	}
	if !hasSpecial {
		return errors.New("password must contain at least one special character")
	}

	return nil
}

func (s *userService) RegisterUser(userRequest *user.UserRegister) (*user.User, error) {

	cek, err := s.repositoryUser.FindByEmail(userRequest.Email)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if cek != nil {
		return nil, errors.New(util.MessageEmailIsNotAvailable)
	}

	err = ValidatePassword(userRequest.Password)
	if err != nil {
		return nil, err
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
