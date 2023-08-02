//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/spf13/viper"

	"elm/internal/handler"
	"elm/internal/middleware"
	"elm/internal/repository"
	"elm/internal/server"
	"elm/internal/service"
	"elm/pkg/helper/sid"
	"elm/pkg/log"
)

var ServerSet = wire.NewSet(server.NewServerHTTP)

var SidSet = wire.NewSet(sid.NewSid)

var JwtSet = wire.NewSet(middleware.NewJwt)

var HandlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
	handler.NewLotteryBallHandler,
)

var ServiceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
	service.NewLotteryBallService,
)

var RepositorySet = wire.NewSet(
	repository.NewDB,
	// repository.NewRedis,
	repository.NewRepository,
	repository.NewUserRepository,
	repository.NewLotteryBallRepository,
)

func newApp(*viper.Viper, *log.Logger) (*gin.Engine, func(), error) {
	panic(wire.Build(
		ServerSet,
		RepositorySet,
		ServiceSet,
		HandlerSet,
		SidSet,
		JwtSet,
	))
}
