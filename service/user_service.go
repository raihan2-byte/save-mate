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
	RegisterUser(userRequest *user.UserRegister) (*user.User, error)
	// FindByUserId(UserId int) (*user.User, error)
	LoginUser(loginRequest *user.UserLoginRequest) (*user.User, error)
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
		return errors.New(util.MessagePasswordMustBeHave6Character)
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
		return errors.New(util.MessagePasswordMustBeHaveUppercase)
	}
	if !hasLower {
		return errors.New(util.MessagePasswordMustBeHaveLowercase)
	}
	if !hasNumber {
		return errors.New(util.MessagePasswordMustBeHaveNumber)
	}
	if !hasSpecial {
		return errors.New(util.MessagePasswordMustBeHaveSpecialCharacter)
	}

	return nil
}

func (s *userService) LoginUser(loginRequest *user.UserLoginRequest) (*user.User, error) {
	err := ValidatePassword(loginRequest.Password)
	if err != nil {
		return nil, err
	}

	cek, err := s.repositoryUser.FindByEmail(loginRequest.Email)
	if err != nil || cek == nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(util.MessageAuthenticationFailed)
		}
		return nil, errors.New(util.MessageAuthenticationFailed)
	}

	err = bcrypt.CompareHashAndPassword([]byte(cek.Password), []byte(loginRequest.Password))
	if err != nil {
		return nil, errors.New(util.MessageAuthenticationFailed)
	}

	return cek, nil
}

func (s *userService) RegisterUser(userRequest *user.UserRegister) (*user.User, error) {

	cek, err := s.repositoryUser.FindByEmail(userRequest.Email)
	if err != nil && err != sql.ErrNoRows {
		return nil, errors.New(util.MessageAuthenticationFailed)
	}

	if cek != nil {
		return nil, errors.New(util.MessageAuthenticationFailed)
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
		return newUser, errors.New(util.MessageAuthenticationFailed)
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
