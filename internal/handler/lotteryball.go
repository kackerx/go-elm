package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"elm/internal/service"
	"elm/pkg/helper/resp"
)

type LotteryBallHandler interface {
	GetLotteryBallById(ctx *gin.Context)
	UpdateLotteryBall(ctx *gin.Context)
	Ping(ctx *gin.Context)
}

type lotteryBallHandler struct {
	*Handler
	lotteryBallService service.LotteryBallService
}

func NewLotteryBallHandler(handler *Handler, lotteryBallService service.LotteryBallService) LotteryBallHandler {
	return &lotteryBallHandler{
		Handler:            handler,
		lotteryBallService: lotteryBallService,
	}
}

func (h *lotteryBallHandler) GetLotteryBallById(ctx *gin.Context) {
	var params struct {
		Id int64 `form:"id" binding:"required"`
	}
	if err := ctx.ShouldBind(&params); err != nil {
		resp.HandleError(ctx, http.StatusBadRequest, 1, err.Error(), nil)
		return
	}

	lotteryBall, err := h.lotteryBallService.GetLotteryBallById(params.Id)
	h.logger.Info("GetLotteryBallByID", zap.Any("lotteryBall", lotteryBall))
	if err != nil {
		resp.HandleError(ctx, http.StatusInternalServerError, 1, err.Error(), nil)
		return
	}
	resp.HandleSuccess(ctx, lotteryBall)
}

func (h *lotteryBallHandler) UpdateLotteryBall(ctx *gin.Context) {
	resp.HandleSuccess(ctx, nil)
}

func (h *lotteryBallHandler) Ping(ctx *gin.Context) {
	resp.HandleSuccess(ctx, map[string]string{"message": "pong"})
}
