package repository

import (
	"elm/internal/model"
)

type Lottery_ballRepository interface {
	FirstById(id int64) (*model.Lottery_ball, error)
}
type lottery_ballRepository struct {
	*Repository
}

func NewLottery_ballRepository(repository *Repository) Lottery_ballRepository {
	return &lottery_ballRepository{
		Repository: repository,
	}
}

func (r *lottery_ballRepository) FirstById(id int64) (*model.Lottery_ball, error) {
	var lottery_ball model.Lottery_ball
	// TODO: query db
	return &lottery_ball, nil
}
