// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"elm/internal/handler"
	"elm/internal/middleware"
	"elm/internal/repository"
	"elm/internal/server"
	"elm/internal/service"
	"elm/pkg/helper/sid"
	"elm/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

// Injectors from wire.go:

func newApp(viperViper *viper.Viper, logger *log.Logger) (*gin.Engine, func(), error) {
	jwt := middleware.NewJwt(viperViper)
	handlerHandler := handler.NewHandler(logger)
	sidSid := sid.NewSid()
	serviceService := service.NewService(logger, sidSid, jwt)
	db := repository.NewDB(viperViper)
	// client := repository.NewRedis(viperViper)
	var client *redis.Client = nil
	repositoryRepository := repository.NewRepository(db, client, logger)
	userRepository := repository.NewUserRepository(repositoryRepository)
	userService := service.NewUserService(serviceService, userRepository)
	userHandler := handler.NewUserHandler(handlerHandler, userService)
	engine := server.NewServerHTTP(logger, jwt, userHandler)
	return engine, func() {
	}, nil
}

// wire.go:

var ServerSet = wire.NewSet(server.NewServerHTTP)

var SidSet = wire.NewSet(sid.NewSid)

var JwtSet = wire.NewSet(middleware.NewJwt)

var HandlerSet = wire.NewSet(handler.NewHandler, handler.NewUserHandler)

var ServiceSet = wire.NewSet(service.NewService, service.NewUserService)

var RepositorySet = wire.NewSet(repository.NewDB, repository.NewRedis, repository.NewRepository, repository.NewUserRepository)