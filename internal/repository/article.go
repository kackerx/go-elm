package repository

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"elm/internal/model"
)

type ArticleRepository interface {
	FirstById(id string) (*model.ArticleContent, error)

	List(offset, pageSize int) ([]*model.Articles, error)

	Create(article *model.Articles) error
}

type articleRepository struct {
	*Repository
}

func (r *articleRepository) List(offset, pageSize int) (res []*model.Articles, err error) {
	if err = r.db.Offset(offset).Limit(pageSize).Find(&res).Order("diy_date desc").Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.Wrap(err, "查找搅珠记录列表失败")
	}

	return
}

func (r *articleRepository) Create(article *model.Articles) error {
	if err := r.db.Create(article).Error; err != nil {
		return errors.Wrap(err, "创建文章失败")
	}

	return nil
}

func NewArticleRepository(repository *Repository) ArticleRepository {
	return &articleRepository{
		Repository: repository,
	}
}

func (r *articleRepository) FirstById(id string) (*model.ArticleContent, error) {
	var article model.ArticleContent
	if err := r.db.Where("id = ?", id).First(&article).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.Wrap(err, "查找文章失败")
	}

	return &article, nil
}
