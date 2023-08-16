package service

import (
	"elm/internal/model"
	"elm/internal/repository"
)

type HomeService interface {
	GetHomeById(id int64) (*model.Home, error)
}

type homeService struct {
	*Service
	homeRepository repository.HomeRepository
}

func NewHomeService(service *Service, homeRepository repository.HomeRepository) HomeService {
	return &homeService{
		Service:        service,
		homeRepository: homeRepository,
	}
}

func (s *homeService) GetHomeById(id int64) (*model.Home, error) {
	return s.homeRepository.FirstById(id)
}
