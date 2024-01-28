package service

import (
	"github.com/deadlorpa/auth-app/interfaces"
	"github.com/deadlorpa/auth-app/repository"
)

type Service struct {
	Authorization interfaces.AuthorizationService
}

func NewService(repositories *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repositories.Users, repositories.Roles),
	}
}
