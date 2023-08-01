package handler

import (
	"github.com/gin-gonic/gin"

	"elm/internal/middleware"
	"elm/pkg/log"
)

type Handler struct {
	logger *log.Logger
}

func NewHandler(logger *log.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}
func GetUserIdFromCtx(ctx *gin.Context) string {
	v, exists := ctx.Get("claims")
	if !exists {
		return ""
	}
	return v.(*middleware.MyCustomClaims).UserId
}
