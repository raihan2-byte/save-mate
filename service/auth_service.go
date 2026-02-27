package service

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserAuthService interface {
	GenerateToken(userID string, role string) (string, error)
	ValidationToken(token string) (*jwt.Token, error)
}

var SECRET_KEY = []byte("test")

type jwtService struct {
}

func NewUserAuthService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID string, role string) (string, error) {
	if userID == "" {
		return "", errors.New("userID cannot be empty")
	}

	if role == "" {
		return "", errors.New("role cannot be empty")
	}

	claim := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(2 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *jwtService) ValidationToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid signing method")
		}
		return SECRET_KEY, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil
}
