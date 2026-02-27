package service

import (
	personaldata "SaveMate/models/PersonalData"
	"SaveMate/models/user"
	"SaveMate/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockPersonalDataRepository struct {
	PersonalData *personaldata.PersonalData
	Err          error
}

func (m *MockPersonalDataRepository) Create(personalData *personaldata.PersonalData) (*personaldata.PersonalData, error) {
	return personalData, nil
}

func (m *MockPersonalDataRepository) FindByPersonalDataId(personalDataId string) (*personaldata.PersonalData, error) {
	return m.PersonalData, nil
}
func (m *MockPersonalDataRepository) FindPersonalDataByUserId(userId string) (*personaldata.PersonalData, error) {
	return m.PersonalData, nil
}

func TestPersonalDataService(t *testing.T) {

	t.Run("TestCreatePersonalData_expectedSuccess", func(t *testing.T) {
		mockPersonalDataRepository := &MockPersonalDataRepository{
			PersonalData: nil,
			Err:          nil,
		}

		mockUserRepository := &MockUserRepository{
			User: &user.User{
				UserId: "12345655",
			},
			Err: nil,
		}

		service := NewPersonalDataService(mockPersonalDataRepository, mockUserRepository)

		request := &personaldata.PersonalDataRequest{
			Sallary:           1000,
			PurposeOfJoinHere: "To Be Rich",
			SavingType:        "FRUGAL",
		}

		result, err := service.CreatePersonalData(request, mockUserRepository.User.UserId)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, result.SavingType, "FRUGAL")
	})

	t.Run("TestCreatePersonalDataWhenUserIdNotFound_expectedFailed", func(t *testing.T) {
		mockPersonalDataRepository := &MockPersonalDataRepository{
			PersonalData: nil,
			Err:          nil,
		}

		mockUserRepository := &MockUserRepository{
			User: nil,
			Err:  nil,
		}

		service := NewPersonalDataService(mockPersonalDataRepository, mockUserRepository)

		request := &personaldata.PersonalDataRequest{
			Sallary:           1000,
			PurposeOfJoinHere: "To Be Rich",
			SavingType:        "FRUGAL",
		}

		result, err := service.CreatePersonalData(request, "")

		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.Equal(t, util.MessageUnauthorized, err.Error())
	})
}
