package repository

import (
	"elm/internal/model"
)

type LotteryBallRepository interface {
	FirstById(id int64) (*model.LotteryBall, error)
}
type lotteryBallRepository struct {
	*Repository
}

func NewLotteryBallRepository(repository *Repository) LotteryBallRepository {
	return &lotteryBallRepository{
		Repository: repository,
	}
}

func (r *lotteryBallRepository) FirstById(id int64) (*model.LotteryBall, error) {
	var lotteryBall model.LotteryBall
	// TODO: query db
	return &lotteryBall, nil
}
