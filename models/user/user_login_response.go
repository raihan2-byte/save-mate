package user

type UserLoginResponse struct {
	Username string
	Email    string
	Token    string
}

func FormatUserLoginResponse(user *User, token string) UserLoginResponse {

	return UserLoginResponse{
		Username: user.Username,
		Email:    user.Email,
		Token:    token,
	}
}
