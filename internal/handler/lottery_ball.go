package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"elm/internal/service"
	"elm/pkg/helper/resp"
)

type Lottery_ballHandler interface {
	GetLottery_ballById(ctx *gin.Context)
	UpdateLottery_ball(ctx *gin.Context)
}

type lottery_ballHandler struct {
	*Handler
	lottery_ballService service.Lottery_ballService
}

func NewLottery_ballHandler(handler *Handler, lottery_ballService service.Lottery_ballService) Lottery_ballHandler {
	return &lottery_ballHandler{
		Handler:             handler,
		lottery_ballService: lottery_ballService,
	}
}

func (h *lottery_ballHandler) GetLottery_ballById(ctx *gin.Context) {
	var params struct {
		Id int64 `form:"id" binding:"required"`
	}
	if err := ctx.ShouldBind(&params); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, err.Error(), nil)
		return
	}

	lottery_ball, err := h.lottery_ballService.GetLottery_ballById(params.Id)
	h.logger.Info("GetLottery_ballByID", zap.Any("lottery_ball", lottery_ball))
	if err != nil {
		resp.HandleError(ctx, http.StatusInternalServerError, 1, err.Error(), nil)
		return
	}
	resp.HandleSuccess(ctx, lottery_ball)
}

func (h *lottery_ballHandler) UpdateLottery_ball(ctx *gin.Context) {
	resp.HandleSuccess(ctx, nil)
}
