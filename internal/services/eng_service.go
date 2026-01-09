package services

import (
	"context"
	"fmt"
	"log"
	"sicbcrg.diamabank.com/internal/models"
	"sicbcrg.diamabank.com/internal/repositories"
)

// EngagementsService contient la logique métier
type EngagementsService struct {
	repo repositories.EngagementsRepository
}

// NewEngagementsService crée une nouvelle instance du service
func NewEngagementsService(repo repositories.EngagementsRepository) *EngagementsService {
	return &EngagementsService{
		repo: repo,
	}
}

// GetAllEngagementsWithDetails récupère tous les engagements avec leurs données associées
func (s *EngagementsService) GetAllEngagementsWithDetails(ctx context.Context) ([]models.Engagements, error) {
	engagements, err := s.repo.GetAll(ctx)
	if err != nil {
		log.Printf("Error fetching engagements: %v", err)
		return nil, err
	}

	// Enrichir chaque engagement avec ses relations
	for i := range engagements {
		if err := s.enrichirEngagement(ctx, &engagements[i]); err != nil {
			log.Printf("Error enriching engagement %s: %v", engagements[i].RefIntEng, err)
			continue
		}
	}

	return engagements, nil
}

// enrichirEngagement charge tous les détails associés à un engagement
func (s *EngagementsService) enrichirEngagement(ctx context.Context, eng *models.Engagements) error {
	// Charger bénéficiaires
	beneficiaires, err := s.repo.GetBeneficiaires(ctx, eng.RefIntEng)
	if err != nil {
		log.Printf("Error loading beneficiaires: %v", err)
	} else if len(beneficiaires) > 0 {
		eng.Beneficiaire = &beneficiaires
	}

	// Charger données additionnelles
	additionnelles, err := s.repo.GetAdditionnelles(ctx, eng.RefIntEng)
	if err != nil {
		log.Printf("Error loading additionnelles: %v", err)
	} else if len(additionnelles) > 0 {
		eng.DonneesAdditionnelles = &additionnelles
	}

	return nil
}

// ExecuteDeclaration exécute la procédure stockée
func (s *EngagementsService) ExecuteDeclaration(ctx context.Context, dateDeclaration string) error {
	if dateDeclaration == "" {
		return fmt.Errorf("date de déclaration is required")
	}

	if err := s.repo.ExecuteStoredProcedure(ctx, dateDeclaration); err != nil {
		log.Printf("Error executing stored procedure: %v", err)
		return err
	}

	return nil
}

// ValidateEngagement valide les données d'un engagement
func (s *EngagementsService) ValidateEngagement(eng *models.Engagements) error {
	if eng.RefIntEng == "" {
		return fmt.Errorf("RefIntEng cannot be empty")
	}
	if eng.TypEng == "" {
		return fmt.Errorf("TypEng cannot be empty")
	}
	return nil
}

// CountEngagements retourne le nombre total d'engagements
func (s *EngagementsService) CountEngagements(ctx context.Context) (int, error) {
	engagements, err := s.repo.GetAll(ctx)
	if err != nil {
		return 0, err
	}
	return len(engagements), nil
}
