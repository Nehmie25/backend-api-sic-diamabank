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
		return "", errors.New("Identifiants invalides")
	}

	if bcrypt.CompareHashAndPassword(
		[]byte(user.MotDePasse),
		[]byte(password),
	) != nil {
		return "", errors.New("Identifiants invalides")
	}

	if user.Isactive == false {
		return "", errors.New("l'utilisateur n'est pas actif. Veuillez contacter l'administrateur pour réactiver votre compte.")
	}

	return s.token.Generate(user.Id, user.Role, user.Email)
}


func (s *AuthService) GetUsers() ([]string, error) {
	users, err := s.userRepo.FindAll()
	if err != nil {
		return nil, errors.New("échec l'ors de la récupération des utilisateurs")
	}

	var userNames []string
	for _, user := range users {
		userNames = append(userNames, user.Nom)
	}

	return userNames, nil
}



