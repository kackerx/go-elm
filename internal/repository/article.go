package repository

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"elm/internal/model"
)

type ArticleRepository interface {
	FirstById(id string) (*model.ArticleContent, error)

	List(offset, pageSize int, cate, year string) ([]*model.Articles, int64, string, error)

	Create(article *model.Articles) error

	CreateContent(article *model.ArticleContent) error

	Update(ctx context.Context, content *model.ArticleContent) error
}

type articleRepository struct {
	*Repository
}

func (r *articleRepository) List(offset, pageSize int, cate, year string) (res []*model.Articles, total int64, yearStr string, err error) {
	switch year {
	case "2021":
		cate = "8"
	case "2020":
		cate = "7"
	case "2019":
		cate = "6"
	case "2018":
		cate = "5"
	case "2017":
		cate = "4"
	case "2016":
		cate = "3"
	case "2022":
		cate = "9"
	case "2023":
		cate = "10"
	case "2024":
		cate = "11"
	case "2025":
		cate = "12"
	default:
		cate = "2"
	}

	r.db.Model(&model.Articles{}).Where("cid = ?", cate).Where("YEAR(STR_TO_DATE(diy_date, '%Y/%m/%d')) = ?", year).Count(&total)

	var years []string
	r.db.Model(&model.Articles{}).
		Select("YEAR(STR_TO_DATE(diy_date, '%Y/%m/%d'))").
		Where("YEAR(STR_TO_DATE(diy_date, '%Y/%m/%d')) is not null").
		Group("YEAR(STR_TO_DATE(diy_date, '%Y/%m/%d'))").
		Order("YEAR(STR_TO_DATE(diy_date, '%Y/%m/%d')) desc").
		Find(&years)
	yearStr = strings.Join(years, ",")

	if err = r.db.
		Where("cid = ?", cate).
		// Where("YEAR(STR_TO_DATE(diy_date, '%Y/%m/%d')) = ?", year).
		Offset(offset).Limit(pageSize).
		Order("diy_qihao desc").
		Find(&res).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, total, yearStr, nil
		}
		return nil, total, yearStr, errors.Wrap(err, "查找搅珠记录列表失败")
	}

	return
}

func (r *articleRepository) Create(article *model.Articles) error {
	var cid string
	switch strconv.Itoa(time.Now().Year()) {
	case "2023":
		cid = "10"
	case "2024":
		cid = "11"
	case "2025":
		cid = "12"
	default:
		cid = "2"
	}
	article.Cid = cid
	if err := r.db.Create(article).Error; err != nil {
		return errors.Wrap(err, "创建文章失败")
	}

	return nil
}

func (r *articleRepository) CreateContent(article *model.ArticleContent) error {
	if err := r.db.Save(article).Error; err != nil {
		return errors.Wrap(err, "创建文章内容失败")
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

func (r *articleRepository) Update(ctx context.Context, content *model.ArticleContent) error {
	if err := r.db.Model(content).Where("id = ?", content.Id).Update("content", content.Content).Error; err != nil {
		return errors.Wrap(err, "更新文章内容失败")
	}

	return nil
}
