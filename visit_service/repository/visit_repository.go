package repository

import (
	"time"
	"visit-service/model"

	"gorm.io/gorm"
)

type VisitRepository struct {
	DB *gorm.DB
}

func NewVisitRepository(db *gorm.DB) *VisitRepository {
	return &VisitRepository{DB: db}
}

func (r *VisitRepository) CreateClient(client *model.Client) error {
	err := r.DB.Create(client).Error
	return err
}

func (r *VisitRepository) CreateAddress(address *model.Address) error {
	err := r.DB.Create(address).Error
	return err
}

func (r *VisitRepository) CreateVisit(visit *model.Visit) error {
	err := r.DB.Create(visit).Error
	return err
}

func (r *VisitRepository) GetVisits() ([]model.Visit, error) {
	var visits []model.Visit

	err := r.DB.
		Preload("Client").
		Preload("Address").
		Where("deleted_at IS NULL").
		Find(&visits).Error

	return visits, err
}

func (r *VisitRepository) DeleteVisit(id uint) error {
	err := r.DB.Table("visits").Where("id = ?", id).Update("deleted_at", time.Now()).Error
	return err
}

func (r *VisitRepository) GetVisitByID(id uint) (model.Visit, error) {
	var visit model.Visit
	err := r.DB.Model(&model.Visit{ID: id}).Scan(&visit).Error
	return visit, err
}

func (r *VisitRepository) UpdateClient(client *model.Client) error {
	return r.DB.Save(client).Error
}

func (r *VisitRepository) UpdateAddress(address *model.Address) error {
	return r.DB.Save(address).Error
}

func (r *VisitRepository) UpdateVisit(visit *model.Visit) error {
	return r.DB.Save(visit).Error
}
