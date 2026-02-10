package repository

import (
	"visit-service/model"

	"gorm.io/gorm"
)

type VisitRepository struct {
	DB *gorm.DB
}

func NewVisitRepository(db *gorm.DB) *VisitRepository {
	return &VisitRepository{DB: db}
}

func (r *VisitRepository) CreateVisit(visit *model.Visit, address *model.Address, client *model.Client) error {
	if err := r.DB.Create(client).Error; err != nil {
		return err
	}

	if err := r.DB.Create(address).Error; err != nil {
		return err
	}

	err := r.DB.Create(visit).Error
	return err
}

func (r *VisitRepository) GetVisits() ([]model.VisitDTO, error) {
	var visits []model.VisitDTO

	err := r.DB.Raw(`
		SELECT v.* FROM visits v
		JOIN addresses ad ON (
			ad.id = v.address_id
		)
		JOIN clients cl ON (
			cl.id = v.client_id
		)
	`).Scan(&visits).Error

	return visits, err
}
