package service

import (
	"elm/internal/model"
	"elm/internal/repository"
)

type Lottery_ballService interface {
	GetLottery_ballById(id int64) (*model.Lottery_ball, error)
}

type lottery_ballService struct {
	*Service
	lottery_ballRepository repository.Lottery_ballRepository
}

func NewLottery_ballService(service *Service, lottery_ballRepository repository.Lottery_ballRepository) Lottery_ballService {
	return &lottery_ballService{
		Service:                service,
		lottery_ballRepository: lottery_ballRepository,
	}
}

func (s *lottery_ballService) GetLottery_ballById(id int64) (*model.Lottery_ball, error) {
	return s.lottery_ballRepository.FirstById(id)
}
