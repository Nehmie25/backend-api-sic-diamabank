package service

import (
	"time"
	"fmt"
	"GoWebapitest/config"
	"GoWebapitest/internal/core/port"
)

type EncoursService struct {
	repo port.EncoursRepository
}

func NewEncoursService(r port.EncoursRepository) *EncoursService {
	return &EncoursService{repo: r}
}

func (s *EncoursService) BuildDeclarationEncours(date string) (config.DeclarationENC, error) {
	// 1. Génération via procédure stockée
	if err := s.repo.GenerateEncours(date); err != nil {
		return config.DeclarationENC{}, err
	}

	// 2. Récupération des données
	encours, err := s.repo.FetchAllEncours()
	if err != nil {
		return config.DeclarationENC{}, err
	}

	// 3. Logique métier
	return config.DeclarationENC{
		NumDec:           fmt.Sprintf("%04d", len(encours)),
		PartEmtr:         "038",
		TypDec:           "01",
		NbrDec:           fmt.Sprintf("%d", len(encours)),
		DateDec:          time.Now().Format("02012006"),
		EncoursEngagement: encours,
	}, nil
}
