package service

import (
	"elm/internal/middleware"
	"elm/pkg/helper/sid"
	"elm/pkg/log"
)

type Service struct {
	logger *log.Logger
	sid    *sid.Sid
	jwt    *middleware.JWT
}

func NewService(logger *log.Logger, sid *sid.Sid, jwt *middleware.JWT) *Service {
	return &Service{
		logger: logger,
		sid:    sid,
		jwt:    jwt,
	}
}
