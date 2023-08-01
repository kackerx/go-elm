package service

import (
	"elm/internal/model"
	"elm/internal/repository"
)

type LotteryBallService interface {
	GetLotteryBallById(id int64) (*model.LotteryBall, error)
}

type lotteryBallService struct {
	*Service
	lotteryBallRepository repository.LotteryBallRepository
}

func NewLotteryBallService(service *Service, lotteryBallRepository repository.LotteryBallRepository) LotteryBallService {
	return &lotteryBallService{
		Service:               service,
		lotteryBallRepository: lotteryBallRepository,
	}
}

func (s *lotteryBallService) GetLotteryBallById(id int64) (*model.LotteryBall, error) {
	return s.lotteryBallRepository.FirstById(id)
}
