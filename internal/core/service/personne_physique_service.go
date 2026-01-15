package service

import (
	"time"
	"fmt"
	"GoWebapitest/config"
	"GoWebapitest/internal/core/port"
)

type PersonnePhysiqueService struct {
	repo port.PersonnePhysiqueRepository
}

func NewPersonnePhysiqueService(r port.PersonnePhysiqueRepository) *PersonnePhysiqueService {
	return &PersonnePhysiqueService{repo: r}
}

func (s *PersonnePhysiqueService) BuildDeclarationepersonnePhysique(date string) (config.DeclarationPP, error) {
	// 1. Génération via procédure stockée
	if err := s.repo.GeneratePersonnesPhysique(date); err != nil {
		return config.DeclarationPP{}, err
	}

	// 2. Récupération des données
	personnes, err := s.repo.FetchAllPersonnesPhysiques()
	if err != nil {
		return config.DeclarationPP{}, err
	}

	// 3. Logique métier
	return config.DeclarationPP{
		NumDec:           fmt.Sprintf("%04d", len(personnes)),
		PartEmtr:         "038",
		TypDec:           "01",
		NbrDec:           fmt.Sprintf("%d", len(personnes)),
		DateDec:          time.Now().Format("02012006"),
		PersonnePhysique: personnes,
	}, nil
}
