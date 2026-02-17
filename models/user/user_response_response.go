package user

import "time"

type UserRegisterResponse struct {
	Id        int
	UserId    string
	Username  string
	Email     string
	Role      string
	CreatedAt time.Time
}
