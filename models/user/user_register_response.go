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

func FormatUserRegisterResponse(user *User) UserRegisterResponse {

	return UserRegisterResponse{
		Id:        user.Id,
		UserId:    user.UserId,
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
	}
}
