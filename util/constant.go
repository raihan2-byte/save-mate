package util

const (
	RoleUser  = "USER"
	RoleAdmin = "ADMIN"

	MessageSuccess              = "success"
	MessageValidationError      = "validation error"
	MessageFailedRegister       = "failed register user"
	MessageEmailIsNotAvailable  = "email already registered"
	MessageAuthenticationFailed = "authentication failed"
	MessageUserNotFound         = "user not found"

	MessagePasswordMustBeHaveUppercase        = "password must contain at least one uppercase letter"
	MessagePasswordMustBeHaveLowercase        = "password must contain at least one lowercase letter"
	MessagePasswordMustBeHaveNumber           = "password must contain at least one number"
	MessagePasswordMustBeHaveSpecialCharacter = "password must contain at least one special character"
	MessagePasswordMustBeHave6Character       = "password must be at least 6 characters"
)
