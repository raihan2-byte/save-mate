package service

import (
	personaldata "SaveMate/models/PersonalData"
	"SaveMate/repository"
	"SaveMate/util"
	"errors"
	"time"

	"github.com/google/uuid"
)

type PersonalDataService interface {
	CreatePersonalData(personalDataRequest *personaldata.PersonalDataRequest) (*personaldata.PersonalData, error)
	GetPersonalDataByUserId(userId string) (*personaldata.PersonalData, error)
}

type personalDataService struct {
	repositoryPersonalData repository.PersonalDataRepository
	repositoryUser         repository.UserRepository
}

func NewPersonalDataService(repositoryPersonalData repository.PersonalDataRepository, repositoryUser repository.UserRepository) *personalDataService {
	return &personalDataService{repositoryPersonalData, repositoryUser}
}

func (s *personalDataService) GetPersonalDataByUserId(userId string) (*personaldata.PersonalData, error) {
	findUserByUserId, err := s.repositoryUser.FindByUserId(userId)
	if err != nil || findUserByUserId == nil {
		return nil, errors.New(util.MessageUnauthorized)
	}

	getData, err := s.repositoryPersonalData.FindPersonalDataByUserId(findUserByUserId.UserId)
	if err != nil {
		return getData, err
	}

	return getData, nil
}

func (s *personalDataService) CreatePersonalData(personalDataRequest *personaldata.PersonalDataRequest, userId string) (*personaldata.PersonalData, error) {

	findUserByUserId, err := s.repositoryUser.FindByUserId(userId)
	if err != nil || findUserByUserId == nil {
		return nil, errors.New(util.MessageUnauthorized)
	}

	generatePersonalDataId := uuid.New().String()
	personalData := &personaldata.PersonalData{
		PersonalDataId:    generatePersonalDataId,
		UserId:            findUserByUserId.UserId,
		Sallary:           personalDataRequest.Sallary,
		PurposeOfJoinHere: personalDataRequest.PurposeOfJoinHere,
		SavingType:        personalDataRequest.SavingType,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	result, err := s.repositoryPersonalData.Create(personalData)
	if err != nil {
		return result, err
	}
	return result, nil
}
