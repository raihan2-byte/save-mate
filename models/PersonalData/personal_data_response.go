package personaldata

import (
	"time"
)

type UserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type PersonalDataResponse struct {
	Id                int
	PersonalDataId    string
	UserId            string
	User              UserResponse
	Sallary           float64
	PurposeOfJoinHere string
	SavingType        string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func FormatPersonalDataResponse(data PersonalData) PersonalDataResponse {

	return PersonalDataResponse{
		Id:             data.Id,
		PersonalDataId: data.PersonalDataId,
		UserId:         data.UserId,
		User: UserResponse{
			Username: data.User.Username,
			Email:    data.User.Email,
		},
		Sallary:           data.Sallary,
		PurposeOfJoinHere: data.PurposeOfJoinHere,
		SavingType:        data.SavingType,
		CreatedAt:         data.CreatedAt,
		UpdatedAt:         data.UpdatedAt,
	}
}
