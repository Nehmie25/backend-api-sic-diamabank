package service

import (
	"time"
	"fmt"
	"GoWebapitest/config"
	"GoWebapitest/internal/core/port"
)

type EngagementsService struct {
	repo port.EngagementRepository
}

func NewEngagementService(r port.EngagementRepository) *EngagementsService {
	return &EngagementsService{repo: r}
}

func (s *EngagementsService) BuildDeclarationEngagements(date string) (config.DeclarationENG, error) {
	// 1. Génération via procédure stockée
	if err := s.repo.GenerateEngagements(date); err != nil {
		return config.DeclarationENG{}, err
	}

	// 2. Récupération des données
	engagements, err := s.repo.FetchAllEngagements()
	if err != nil {
		return config.DeclarationENG{}, err
	}

	// 3. Logique métier
	return config.DeclarationENG{
		NumDec:           fmt.Sprintf("%04d", len(engagements)),
		PartEmtr:         "038",
		TypDec:           "01",
		NbrDec:           fmt.Sprintf("%d", len(engagements)),
		DateDec:          time.Now().Format("02012006"),
		Engagement: engagements,
	}, nil
}
