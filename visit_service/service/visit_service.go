package service

import (
	"fmt"
	"log"
	"strconv"
	"visit-service/client"
	"visit-service/config"
	"visit-service/model"
	"visit-service/repository"
)

type VisitService struct {
	repository *repository.VisitRepository
	client     client.UserServiceClient
}

func NewVisitService(repository *repository.VisitRepository) *VisitService {
	cfg := config.GetConfig()
	url := fmt.Sprintf("http://%s:%s", cfg.UserServiceHost, cfg.UserServicePort)
	return &VisitService{repository: repository, client: client.NewUserServiceClient(url)}
}

func (s *VisitService) CreateVisit(visit *model.VisitDTO) error {
	client := &model.Client{
		FirstName: visit.Client.FirstName,
		LastName:  visit.Client.LastName,
		Surname:   visit.Client.Surname,
		Phone:     visit.Client.Phone,
	}

	if err := s.repository.CreateClient(client); err != nil {
		return fmt.Errorf("failed to create client: %w", err)
	}

	address := &model.Address{
		City:             visit.Address.City,
		Locality:         visit.Address.Locality,
		Region:           visit.Address.Region,
		Street:           visit.Address.Street,
		HouseNumber:      visit.Address.HouseNumber,
		Letter:           visit.Address.Letter,
		Building:         visit.Address.Building,
		AppartmentNumber: visit.Address.AppartmentNumber,
	}

	if err := s.repository.CreateAddress(address); err != nil {
		return fmt.Errorf("failed to create address: %w", err)
	}

	visitDB := &model.Visit{
		MasterID:             visit.MasterID,
		ContractNumber:       visit.ContractNumber,
		ContractDate:         visit.ContractDate,
		ScheduledDate:        visit.ScheduledDate,
		ScheduledTime:        visit.ScheduledTime,
		EquipmentDescription: visit.EquipmentDescription,
		AssignedMonth:        visit.AssignedMonth,
		Amount:               visit.Amount,

		Client:  *client,
		Address: *address,
	}

	err := s.repository.CreateVisit(visitDB)
	if err != nil {
		return err
	}

	return nil
}

func (s *VisitService) GetVisits() ([]model.VisitDTO, error) {
	visits, err := s.repository.GetVisits()
	if err != nil {
		log.Printf("error getting visits from DB: %v", err)
		return []model.VisitDTO{}, nil
	}

	var masterIDs []uint
	for _, vis := range visits {
		if vis.MasterID != nil {
			masterIDs = append(masterIDs, *vis.MasterID)
		}
	}

	var masters []model.MasterDTO
	if len(masterIDs) > 0 {
		masters, err = s.client.GetMastersByIDs(masterIDs)
		if err != nil {
			return nil, err
		}
	}

	mastersMap := make(map[uint]model.MasterDTO)
	for _, master := range masters {
		mastersMap[master.ID] = master
	}

	var visitsDTO []model.VisitDTO
	for _, vis := range visits {
		visDTO := model.VisitDTO{
			Client: struct {
				FirstName string `json:"first_name"`
				LastName  string `json:"last_name"`
				Surname   string `json:"surname"`
				Phone     string `json:"phone"`
			}{
				FirstName: vis.Client.FirstName,
				LastName:  vis.Client.LastName,
				Surname:   vis.Client.Surname,
				Phone:     vis.Client.Phone,
			},
			Address: struct {
				City             string `json:"city"`
				Locality         string `json:"locality"`
				Region           string `json:"region"`
				Street           string `json:"street"`
				HouseNumber      uint   `json:"house_number"`
				Letter           string `json:"letter"`
				Building         *uint  `json:"building"`
				AppartmentNumber *uint  `json:"appartment_number"`
			}{
				City:             vis.Address.City,
				Locality:         vis.Address.Locality,
				Region:           vis.Address.Region,
				Street:           vis.Address.Street,
				HouseNumber:      vis.Address.HouseNumber,
				Letter:           vis.Address.Letter,
				Building:         vis.Address.Building,
				AppartmentNumber: vis.Address.AppartmentNumber,
			},
			ID:                   vis.ID,
			ContractNumber:       vis.ContractNumber,
			ContractDate:         vis.ContractDate,
			ScheduledDate:        vis.ScheduledDate,
			ScheduledTime:        vis.ScheduledTime,
			EquipmentDescription: vis.EquipmentDescription,
			AssignedMonth:        vis.AssignedMonth,
			Amount:               vis.Amount,
			MasterID:             vis.MasterID,
		}

		visitsDTO = append(visitsDTO, visDTO)
	}

	return visitsDTO, nil
}

func (s *VisitService) DeleteVisit(idString string) error {
	id, err := strconv.Atoi(idString)
	if err != nil {
		return err
	}

	err = s.repository.DeleteVisit(uint(id))

	return err
}
