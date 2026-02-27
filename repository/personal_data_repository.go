package repository

import (
	personaldata "SaveMate/models/PersonalData"
	"database/sql"
)

type PersonalDataRepository interface {
	Create(personalData *personaldata.PersonalData) (*personaldata.PersonalData, error)
	FindByPersonalDataId(personalDataId string) (*personaldata.PersonalData, error)
	FindPersonalDataByUserId(userId string) (*personaldata.PersonalData, error)
}

type personalDataRepository struct {
	db *sql.DB
}

func NewPersonalDataRepository(db *sql.DB) *personalDataRepository {
	return &personalDataRepository{db}
}

func (r *personalDataRepository) Create(personalData *personaldata.PersonalData) (*personaldata.PersonalData, error) {
	query := "INSERT INTO personal_data (id, personal_data_id, user_id, sallary, purpose_of_join_here, saving_type, created_at, updated_at) VALUES (?,?,?,?,?,?,?,?)"

	result, err := r.db.Exec(query, personalData.Id, personalData.PersonalDataId, personalData.UserId, personalData.Sallary, personalData.PurposeOfJoinHere, personalData.SavingType, personalData.CreatedAt, personalData.UpdatedAt)
	if err != nil {
		return personalData, err
	}

	idPersonalData, _ := result.LastInsertId()
	personalData.Id = int(idPersonalData)

	return personalData, nil
}
