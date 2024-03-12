package service

import (
	"medods"
	"medods/pkg/repository"
)

type Token interface {
	SetRefreshToken(guid string, refresh []byte) error
	CheckUser(guid string) error
	GetUser(refresh string) (medods.User, error)
	UpdateTokens(user medods.User) error
}

type Service struct {
	Token
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Token: NewTokenService(repos.Token),
	}
}
