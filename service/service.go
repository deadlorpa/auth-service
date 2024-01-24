package service

import "github.com/deadlorpa/auth-app/repository"

type Authorization interface {
}

type Service struct {
	Authorization
}

func NewService(repositories *repository.Repository) *Service {
	return &Service{}
}
