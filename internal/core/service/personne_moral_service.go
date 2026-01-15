package service

import (
	"time"
	"fmt"
	"GoWebapitest/config"
	"GoWebapitest/internal/core/port"
)

type PersonneMoraleService struct {
	repo port.PersonneMoralRepository
}

func NewPersonneMoralService(r port.PersonneMoralRepository) *PersonneMoraleService {
	return &PersonneMoraleService{repo: r}
}

func (s *PersonneMoraleService) BuildDeclarationPersonneMorale(date string) (config.DeclarationPM, error) {
	// 1. Génération via procédure stockée
	if err := s.repo.GeneratePersonnesMoral(date); err != nil {
		return config.DeclarationPM{}, err
	}

	// 2. Récupération des données
	personnes, err := s.repo.FetchAllPersonnesMoral()
	if err != nil {
		return config.DeclarationPM{}, err
	}

	// 3. Logique métier
	return config.DeclarationPM{
		NumDec:           fmt.Sprintf("%04d", len(personnes)),
		PartEmtr:         "038",
		TypDec:           "02",
		NbrDec:           fmt.Sprintf("%d", len(personnes)),
		DateDec:          time.Now().Format("02012006"),
		PersonneMorale: personnes,
	}, nil
}
