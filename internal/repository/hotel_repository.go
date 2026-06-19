package repository

import (
	"hotel-go/internal/model"

	"gorm.io/gorm"
)

type HotelRepository struct {
	db *gorm.DB
}

func NewHotelRepository(db *gorm.DB) *HotelRepository {
	return &HotelRepository{
		db: db,
	}
}

func (r *HotelRepository) FindAll() ([]model.Hotel, error) {
	var hotels []model.Hotel

	err := r.db.
		Preload("Standort").
		Preload("Zimmer").
		Find(&hotels).Error

	if err != nil {
		return nil, err
	}

	return hotels, nil
}

func (r *HotelRepository) FindByID(id uint) (*model.Hotel, error) {
	var hotel model.Hotel

	err := r.db.
		Preload("Standort").
		Preload("Zimmer").
		First(&hotel, id).Error

	if err != nil {
		return nil, err
	}

	return &hotel, nil
}

func (r *HotelRepository) Create(hotel *model.Hotel) error {
	return r.db.Create(hotel).Error
}
