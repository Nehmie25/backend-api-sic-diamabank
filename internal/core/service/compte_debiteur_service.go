package service

import (
	"time"
	"fmt"
	"GoWebapitest/config"
	"GoWebapitest/internal/core/port"
)

type CompteDebiteursService struct {
	repo port.CompteDebiteursRepository
}

func NewCompteDebiteursService(r port.CompteDebiteursRepository) *CompteDebiteursService {
	return &CompteDebiteursService{repo: r}
}

func (s *CompteDebiteursService) BuildDeclarationEncours(date string) (config.DeclarationDEB, error) {
	// 1. Génération via procédure stockée
	if err := s.repo.GenerateEncours(date); err != nil {
		return config.DeclarationDEB{}, err
	}

	// 2. Récupération des données
	encours, err := s.repo.FetchAllEncours()
	if err != nil {
		return config.DeclarationDEB{}, err
	}

	// 3. Logique métier
	return config.DeclarationDEB{
		NumDec:           fmt.Sprintf("%04d", len(encours)),
		PartEmtr:         "038",
		TypDec:           "01",
		NbrDec:           fmt.Sprintf("%d", len(encours)),
		DateDec:          time.Now().Format("02012006"),
		CompteDebiteurs: encours,
	}, nil
}
