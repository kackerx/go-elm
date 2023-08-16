package repository

import (
	"elm/internal/model"
)

type HomeRepository interface {
	FirstById(id int64) (*model.Home, error)
}
type homeRepository struct {
	*Repository
}

func NewHomeRepository(repository *Repository) HomeRepository {
	return &homeRepository{
		Repository: repository,
	}
}

func (r *homeRepository) FirstById(id int64) (*model.Home, error) {
	var home model.Home
	// TODO: query db
	return &home, nil
}
