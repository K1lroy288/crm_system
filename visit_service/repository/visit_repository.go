package repository

import "gorm.io/gorm"

type VisitRepository struct {
	DB *gorm.DB
}

func NewVisitRepository(db *gorm.DB) *VisitRepository {
	return &VisitRepository{DB: db}
}
