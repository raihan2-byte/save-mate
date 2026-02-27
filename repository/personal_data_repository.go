package repository

import (
	personaldata "SaveMate/models/PersonalData"
	"database/sql"
)

type PersonalDataRepository interface {
	Create(personalData *personaldata.PersonalData) (*personaldata.PersonalData, error)
	FindByPersonalDataId(personalDataId string) (*personaldata.PersonalData, error)
	FindPersonalDataByUserId(userId string) (*personaldata.PersonalData, error)
	FindAllPersonalData(limit int, offset int) ([]*personaldata.PersonalData, error)
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

func (r *personalDataRepository) FindPersonalDataByUserId(userId string) (*personaldata.PersonalData, error) {
	query := "SELECT pd.personal_data_id, pd.user_id, pd.sallary, pd.purpose_of_join_here, pd.saving_type, pd.created_at, pd.updated_at, u.user_id, u.username u.email from personal_datas pd JOIN users u ON u.user_id = pd.user_id WHERE pd.user_id = ? "

	row := r.db.QueryRow(query, userId)

	personalData := &personaldata.PersonalData{}
	err := row.Scan(&personalData.PersonalDataId, &personalData.UserId, &personalData.Sallary,
		&personalData.PurposeOfJoinHere, &personalData.SavingType, &personalData.CreatedAt, personalData.UpdatedAt, &personalData.User.UserId, &personalData.User.Username,
		&personalData.User.Email)

	if err != nil {
		return personalData, err
	}
	return personalData, nil
}

func (r *personalDataRepository) FindAllPersonalData(limit int, offset int) ([]*personaldata.PersonalData, error) {
	query := "SELECT pd.personal_data_id, pd.user_id, pd.sallary, pd.purpose_of_join_here, pd.saving_type, pd.created_at, pd.updated_at, u.user_id, u.username, u.email from personal_datas pd JOIN users u ON  u.user_id = pd.user_id ORDER BY pd.created_at DESC LIMIT ? OFFSET ?"

	row, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}

	var personalDatas []*personaldata.PersonalData

	for row.Next() {

		pd := &personaldata.PersonalData{}
		err := row.Scan(&pd.PersonalDataId, &pd.UserId, &pd.Sallary,
			&pd.PurposeOfJoinHere, &pd.SavingType, &pd.CreatedAt, pd.UpdatedAt, &pd.User.UserId, &pd.User.Username,
			&pd.User.Email)

		if err != nil {
			return nil, err
		}
		personalDatas = append(personalDatas, pd)
	}
	if err = row.Err(); err != nil {
		return nil, err
	}

	return personalDatas, nil
}

func (r *personalDataRepository) FindByPersonalDataId(personalDataId string) (*personaldata.PersonalData, error) {
	query := "SELECT pd.personal_data_id, pd.user_id, pd.sallary, pd.purpose_of_join_here, pd.saving_type, pd.created_at, pd.updated_at, u.user_id, u.username, u.email from personal_datas pd JOIN user u ON pd.user_id = u.user_id ORDER BY pd.created_at DESC WHERE personal_data_id = ?"

	row := r.db.QueryRow(query, personalDataId)

	personalData := &personaldata.PersonalData{}
	err := row.Scan(&personalData.PersonalDataId, &personalData.UserId, &personalData.Sallary,
		&personalData.PurposeOfJoinHere, &personalData.SavingType, &personalData.CreatedAt, personalData.UpdatedAt, &personalData.User.UserId, &personalData.User.Username,
		&personalData.User.Email)

	if err != nil {
		return personalData, err
	}
	return personalData, nil
}
