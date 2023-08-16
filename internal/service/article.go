package service

import (
	"time"

	"elm/internal/model"
	"elm/internal/repository"
	"elm/utils"
	"elm/vars"
)

type AddArticleRequest struct {
	Title     string `json:"title" binding:"required"`
	Qihao     string `json:"qihao" binding:"required"`
	Data      string `json:"data" binding:"required"`
	Shengxiao string `json:"shengxiao" binding:"required"`
	Tema      string `json:"tema" binding:"required"`
}

type ArticleService interface {
	GetArticleById(string) (*model.ArticleContent, error)

	GetArticleList(vars.PageParams) ([]*model.Articles, error)

	AddArticle(request *AddArticleRequest) error
}

type articleService struct {
	*Service
	articleRepository repository.ArticleRepository
}

func (s *articleService) AddArticle(req *AddArticleRequest) error {
	return s.articleRepository.Create(&model.Articles{
		Title:        req.Title,
		DiyDate:      time.Now().Format(`2006/01/02`),
		DiyQihao:     req.Qihao,
		DiyData:      req.Data,
		DiyShengxiao: req.Shengxiao,
		DiyTema:      req.Tema,
	})
}

func (s *articleService) GetArticleList(params vars.PageParams) ([]*model.Articles, error) {
	offset := utils.GetPageOffset(params.PageSize, params.PageNum)

	return s.articleRepository.List(offset, params.PageSize)
}

func NewArticleService(service *Service, articleRepository repository.ArticleRepository) ArticleService {
	return &articleService{
		Service:           service,
		articleRepository: articleRepository,
	}
}

func (s *articleService) GetArticleById(id string) (*model.ArticleContent, error) {
	return s.articleRepository.FirstById(id)
}
