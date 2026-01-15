package service

import (
	"GoWebapitest/internal/core/port"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo port.UserRepository
	token    port.TokenGenerator
}

func NewAuthService(
	repo port.UserRepository,
	token port.TokenGenerator,
) *AuthService {
	return &AuthService{
		userRepo: repo,
		token:    token,
	}
}

func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.userRepo.FindByEmail(email)

	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if bcrypt.CompareHashAndPassword(
		[]byte(user.MotDePasse),
		[]byte(password),
	) != nil {
		return "", errors.New("invalid credentials")
	}

	return s.token.Generate(user.Id, user.Role)
}
