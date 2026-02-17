package user

import "time"

type User struct {
	Id        int
	UserId    string
	Username  string
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
