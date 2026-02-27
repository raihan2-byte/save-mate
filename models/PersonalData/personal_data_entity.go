package personaldata

import (
	"SaveMate/models/user"
	"time"
)

type PersonalData struct {
	Id                int
	PersonalDataId    string
	UserId            string
	User              user.User
	Sallary           float64
	PurposeOfJoinHere string
	SavingType        string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
