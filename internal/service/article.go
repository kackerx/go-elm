package service

import (
	"context"
	"time"

	"elm/internal/model"
	"elm/internal/repository"
	"elm/utils"
	"elm/vars"
)

type AddArticleRequest struct {
	Qihao        string `json:"qihao" binding:"required"`
	DiyData      string `json:"diy_data" binding:"required"`
	DiyShengxiao string `json:"diy_shengxiao" binding:"required"`
	DiyTema      string `json:"diy_tema" binding:"required"`
}

type AddArticleContentRequest struct {
	Id      string `json:"id" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type ArticleService interface {
	GetArticleById(string) (*model.ArticleContent, error)

	GetArticleList(vars.PageParams, string, string) ([]*model.Articles, int64, string, error)

	GetArticleToday(ctx context.Context) (*model.Articles, bool, error)

	AddArticle(request *AddArticleRequest) error

	AddArticleContent(request *AddArticleContentRequest) error

	SaveArticleImg(imgName, title string) error

	UpdateArticleContent(ctx context.Context, request *AddArticleContentRequest) error

	DeleteArticle(ctx context.Context, id string) error

	GetImgList(vars.PageParams) ([]*model.Articles, int64, error)
}

type articleService struct {
	*Service
	articleRepository repository.ArticleRepository
}

func (s *articleService) AddArticle(req *AddArticleRequest) error {
	m := &model.Articles{
		Title:        time.Now().Format("2006/1/02") + "---" + req.Qihao + "开奖结果",
		DiyDate:      time.Now().Format("2006/1/02"),
		DiyQihao:     req.Qihao,
		DiyData:      req.DiyData,
		DiyShengxiao: req.DiyShengxiao,
		DiyTema:      req.DiyTema,
		PublishTime:  time.Now().Unix(),
	}

	return s.articleRepository.Create(m)
}

func (s *articleService) AddArticleContent(req *AddArticleContentRequest) error {
	return s.articleRepository.CreateContent(&model.ArticleContent{
		Id:      req.Id,
		Content: req.Content,
	})
}

func (s *articleService) GetArticleList(params vars.PageParams, cate, year string) ([]*model.Articles, int64, string, error) {
	offset := utils.GetPageOffset(params.PageSize, params.PageNum)

	return s.articleRepository.List(offset, params.PageSize, cate, year, params.IsAdmin)
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

func (s *articleService) SaveArticleImg(imgName, title string) error {
	return s.articleRepository.Create(&model.Articles{
		Title:       title,
		PublishTime: time.Now().Unix(),
		ImgUrl:      imgName,
		Cid:         "1",
		DiyQihao:    "0",
		DiyTema:     "0",
	})
}

func (s *articleService) UpdateArticleContent(ctx context.Context, request *AddArticleContentRequest) error {
	return s.articleRepository.Update(ctx, &model.ArticleContent{
		Id:      request.Id,
		Content: request.Content,
	})
}

func (s *articleService) DeleteArticle(ctx context.Context, id string) error {
	return s.articleRepository.Delete(ctx, id)
}

func (s *articleService) GetArticleToday(ctx context.Context) (*model.Articles, bool, error) {
	return s.articleRepository.First(ctx)
}

func (s *articleService) GetImgList(params vars.PageParams) ([]*model.Articles, int64, error) {
	offset := utils.GetPageOffset(params.PageSize, params.PageNum)

	return s.articleRepository.ListImg(offset, params.PageSize)

}
