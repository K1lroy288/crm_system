package service

import (
	"visit-service/client"
	"visit-service/model"
	"visit-service/repository"
)

type VisitService struct {
	repository *repository.VisitRepository
	client     client.UserServiceClient
}

func NewVisitService(repository *repository.VisitRepository) *VisitService {
	return &VisitService{repository: repository}
}

func (s *VisitService) CreateVisit(visit *model.VisitDTO) error {
	return nil
}

func (s *VisitService) GetVisits() ([]model.VisitDTO, error) {
	visits, err := s.repository.GetVisits()
	if err != nil {
		return nil, err
	}

	var masterIDs []uint
	for _, vis := range visits {
		masterIDs = append(masterIDs, *vis.MasterID)
	}

	masters, err := s.client.GetMastersByIDs(masterIDs)
	if err != nil {
		return nil, err
	}

	var visitsDTO []model.VisitDTO
	for _, vis := range visits {
		currMaster := &model.MasterDTO{}
		for _, master := range masters {
			if vis.MasterID == &master.ID {
				currMaster = &master
			}
		}

		visDTO := &model.VisitDTO{
			FirstName:            vis.Client.FirstName,
			LastName:             vis.Client.LastName,
			Surname:              vis.Client.Surname,
			Phone:                vis.Client.Phone,
			MasterID:             vis.MasterID,
			MasterUsername:       currMaster.Username,
			MasterFirstname:      currMaster.FirstName,
			MasterLastname:       currMaster.LastName,
			City:                 vis.Address.City,
			Locality:             vis.Address.Locality,
			Region:               vis.Address.Region,
			Street:               vis.Address.Street,
			HouseNumber:          vis.Address.HouseNumber,
			Letter:               vis.Address.Letter,
			Building:             vis.Address.Building,
			Appartment_number:    vis.Address.Appartment_number,
			ContractNumber:       vis.ContractNumber,
			ContractDate:         vis.ContractDate,
			ScheduledDate:        vis.ScheduledDate,
			ScheduledTime:        vis.ScheduledTime,
			EquipmentDescription: vis.EquipmentDescription,
			AssignedMonth:        vis.AssignedMonth,
			Amount:               vis.Amount,
		}

		visitsDTO = append(visitsDTO, *visDTO)
	}

	return visitsDTO, nil
}
