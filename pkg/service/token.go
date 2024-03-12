package service

import (
	"medods"
	"medods/pkg/repository"
)

type TokenService struct {
	repo repository.Token
}

func NewTokenService(repo repository.Token) *TokenService {
	return &TokenService{repo: repo}
}

func (s *TokenService) SetRefreshToken(guid string, refresh []byte) error {
	return s.repo.SetRefreshToken(guid, refresh)
}

func (s *TokenService) CheckUser(guid string) error {
	return s.repo.CheckUser(guid)
}

func (s *TokenService) GetUser(refresh string) (medods.User, error) {
	return s.repo.GetUser(refresh)
}

func (s *TokenService) UpdateTokens(user medods.User) error {
	return s.repo.UpdateTokens(user)
}
