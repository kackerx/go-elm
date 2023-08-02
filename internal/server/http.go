package server

import (
	"github.com/gin-gonic/gin"

	"elm/internal/handler"
	"elm/internal/middleware"
	"elm/pkg/helper/resp"
	"elm/pkg/log"
)

func NewServerHTTP(
	logger *log.Logger,
	jwt *middleware.JWT,
	userHandler handler.UserHandler,
	lotteryBallHandler handler.LotteryBallHandler,
) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(
		middleware.CORSMiddleware(),
		middleware.ResponseLogMiddleware(logger),
		middleware.RequestLogMiddleware(logger),
		// middleware.SignMiddleware(log),
	)

	// No route group has permission
	noAuthRouter := r.Group("/")
	{

		noAuthRouter.GET("/", func(ctx *gin.Context) {
			logger.WithContext(ctx).Info("hello")
			resp.HandleSuccess(ctx, map[string]interface{}{
				"say": "Hi Nunu!",
			})
		})

		noAuthRouter.POST("/register", userHandler.Register)
		noAuthRouter.POST("/login", userHandler.Login)
	}
	// Non-strict permission routing group
	noStrictAuthRouter := r.Group("/").Use(middleware.NoStrictAuth(jwt, logger))
	{
		noStrictAuthRouter.GET("/user", userHandler.GetProfile)
	}

	// Strict permission routing group
	strictAuthRouter := r.Group("/").Use(middleware.StrictAuth(jwt, logger))
	{
		strictAuthRouter.PUT("/user", userHandler.UpdateProfile)
	}

	// vue-test api
	vueTestRouter := r.Group("/api")
	{
		vueTestRouter.GET("/", lotteryBallHandler.Ping)
	}

	return r
}
