package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"elm/internal/service"
	"elm/pkg/helper/resp"
	"elm/vars"
)

type ArticleHandler interface {
	GetArticleById(ctx *gin.Context)
	UpdateArticle(ctx *gin.Context)

	GetArticleList(ctx *gin.Context)

	AddArticle(ctx *gin.Context)
}

type articleHandler struct {
	*Handler
	articleService service.ArticleService
}

func NewArticleHandler(handler *Handler, articleService service.ArticleService) ArticleHandler {
	return &articleHandler{
		Handler:        handler,
		articleService: articleService,
	}
}

func (h *articleHandler) AddArticle(ctx *gin.Context) {
	req := new(service.AddArticleRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, err.Error(), nil)
		return
	}

	err := h.articleService.AddArticle(req)
	if err != nil {
		resp.HandleError(ctx, http.StatusOK, 1, err.Error(), nil)
		return
	}
	resp.HandleSuccess(ctx, nil)
}

func (h *articleHandler) GetArticleList(ctx *gin.Context) {
	var params vars.PageParams
	if err := ctx.ShouldBind(&params); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, err.Error(), nil)
		return
	}

	articles, err := h.articleService.GetArticleList(params)
	if err != nil {
		resp.HandleError(ctx, http.StatusOK, 1, err.Error(), nil)
		return
	}
	resp.HandleSuccess(ctx, articles)
}

func (h *articleHandler) GetArticleById(ctx *gin.Context) {
	// var params struct {
	// 	Id int64 `form:"id" binding:"required"`
	// }
	// if err := ctx.ShouldBind(&params); err != nil {
	// 	resp.HandleError(ctx, http.StatusBadRequest, 1, err.Error(), nil)
	// 	return
	// }

	param := ctx.Param("id")
	if param == "" {
		resp.HandleError(ctx, http.StatusBadRequest, 1, "文章id为空", nil)
		return
	}

	article, err := h.articleService.GetArticleById(param)
	h.logger.Info("GetArticleByID", zap.Any("article", article))
	if err != nil {
		resp.HandleError(ctx, http.StatusInternalServerError, 1, err.Error(), nil)
		return
	}
	resp.HandleSuccess(ctx, article)
}

func (h *articleHandler) UpdateArticle(ctx *gin.Context) {
	resp.HandleSuccess(ctx, nil)
}
