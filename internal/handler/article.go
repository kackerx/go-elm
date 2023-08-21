package handler

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"elm/internal/service"
	"elm/pkg/helper/resp"
	"elm/utils"
	"elm/vars"
)

type ArticleHandler interface {
	GetArticleById(ctx *gin.Context)

	UpdateArticle(ctx *gin.Context)

	GetArticleList(ctx *gin.Context)

	AddArticle(ctx *gin.Context)

	AddArticleContent(ctx *gin.Context)

	UpdateArticleContent(ctx *gin.Context)

	ImageUpload(ctx *gin.Context)
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

func (h *articleHandler) AddArticleContent(ctx *gin.Context) {
	req := new(service.AddArticleContentRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, err.Error(), nil)
		return
	}

	err := h.articleService.AddArticleContent(req)
	if err != nil {
		resp.HandleError(ctx, http.StatusOK, 1, err.Error(), nil)
		return
	}
	resp.HandleSuccess(ctx, nil)
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

	articles, err := h.articleService.GetArticleList(params, ctx.Query("cate"), ctx.Query("year"))
	if err != nil {
		resp.HandleError(ctx, http.StatusOK, 1, "获取记录失败", nil)
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

func (h *articleHandler) UpdateArticleContent(ctx *gin.Context) {
	resp.HandleSuccess(ctx, nil)
}

func (h *articleHandler) ImageUpload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		resp.HandleError(ctx, http.StatusOK, vars.ErrorMap[vars.ErrImgUploadFail].Code, vars.ErrorMap[vars.ErrImgUploadFail].Error(), nil)
		return
	}

	fileName := utils.GetImagName(file.Filename, filepath.Ext(file.Filename))
	err = ctx.SaveUploadedFile(file, fileName)
	if err != nil {
		resp.HandleError(ctx, http.StatusOK, vars.ErrorMap[vars.ErrImgUploadFail].Code, vars.ErrorMap[vars.ErrImgUploadFail].Error(), nil)
		return
	}

	err = h.articleService.SaveArticleImg("/"+fileName, ctx.PostForm("title"))
	if err != nil {
		resp.HandleError(ctx, http.StatusOK, vars.ErrorMap[vars.ErrImgUploadFail].Code, vars.ErrorMap[vars.ErrImgUploadFail].Error(), nil)
		return
	}

	resp.HandleSuccess(ctx, nil)
}
