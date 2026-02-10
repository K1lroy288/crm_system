package service

import (
	"visit-service/model"
	"visit-service/repository"
)

type VisitService struct {
	repository *repository.VisitRepository
}

func NewVisitService(repository *repository.VisitRepository) *VisitService {
	return &VisitService{repository: repository}
}

func (s *VisitService) CreateVisit(visit *model.VisitDTO) error {
	return nil
}
