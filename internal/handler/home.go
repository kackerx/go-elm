package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"elm/internal/service"
	"elm/pkg/helper/resp"
)

type HomeHandler interface {
	GetHomeById(ctx *gin.Context)
	UpdateHome(ctx *gin.Context)
	GetHomePage(ctx *gin.Context)
}

type homeHandler struct {
	*Handler
	homeService service.HomeService
}

func (h *homeHandler) GetHomePage(ctx *gin.Context) {
	resp.HandleSuccess(ctx, map[string]interface{}{
		"banner": map[string]string{
			"imgUrl": "hehe",
		},
	})
}

func NewHomeHandler(handler *Handler, homeService service.HomeService) HomeHandler {
	return &homeHandler{
		Handler:     handler,
		homeService: homeService,
	}
}

func (h *homeHandler) GetHomeById(ctx *gin.Context) {
	var params struct {
		Id int64 `form:"id" binding:"required"`
	}
	if err := ctx.ShouldBind(&params); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, err.Error(), nil)
		return
	}

	home, err := h.homeService.GetHomeById(params.Id)
	h.logger.Info("GetHomeByID", zap.Any("home", home))
	if err != nil {
		resp.HandleError(ctx, http.StatusInternalServerError, 1, err.Error(), nil)
		return
	}
	resp.HandleSuccess(ctx, home)
}

func (h *homeHandler) UpdateHome(ctx *gin.Context) {
	resp.HandleSuccess(ctx, nil)
}
