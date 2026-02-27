package personaldata

type PersonalDataRequest struct {
	Sallary           float64 `json:"sallary" binding:"required"`
	PurposeOfJoinHere string  `json:"purpose_of_join_here" binding:"required"`
	SavingType        string  `json:"saving_type" binding:"required"`
}
