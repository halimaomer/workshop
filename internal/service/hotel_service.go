package service

import (
	"hotel-go/internal/model"
	"hotel-go/internal/repository"
)

type HotelService struct {
	repo *repository.HotelRepository
}

func NewHotelService(repo *repository.HotelRepository) *HotelService {
	return &HotelService{
		repo: repo,
	}
}

func (s *HotelService) FindAll() ([]model.Hotel, error) {
	return s.repo.FindAll()
}

func (s *HotelService) FindByID(id uint) (*model.Hotel, error) {
	return s.repo.FindByID(id)
}

func (s *HotelService) Create(hotel *model.Hotel) error {
	return s.repo.Create(hotel)
}
