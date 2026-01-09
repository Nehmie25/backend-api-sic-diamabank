package services

import (
	"context"
	"fmt"
	"log"
	"sicbcrg.diamabank.com/internal/models"
	"sicbcrg.diamabank.com/internal/repositories"
)

// EncoursService contient la logique métier
type EncoursService struct {
	repo repositories.EncoursRepository
}

// NewEncoursService crée une nouvelle instance du service
func NewEncoursService(repo repositories.EncoursRepository) *EncoursService {
	return &EncoursService{
		repo: repo,
	}
}

// GetAllEncoursWithDetails récupère tous les encours avec leurs données associées
func (s *EncoursService) GetAllEncoursWithDetails(ctx context.Context) ([]models.Encours, error) {
	encours, err := s.repo.GetAll(ctx)
	if err != nil {
		log.Printf("Error fetching encours: %v", err)
		return nil, err
	}

	// Enrichir chaque encours avec ses relations
	for i := range encours {
		if err := s.enrichirEncours(ctx, &encours[i]); err != nil {
			log.Printf("Error enriching encours %s: %v", encours[i].RefIntEng, err)
			continue
		}
	}

	return encours, nil
}

// enrichirEncours charge tous les détails associés à un encours
func (s *EncoursService) enrichirEncours(ctx context.Context, enc *models.Encours) error {
	// Charger données additionnelles
	additionnelles, err := s.repo.GetAdditionnelles(ctx, enc.RefIntEng)
	if err != nil {
		log.Printf("Error loading additionnelles: %v", err)
	} else if len(additionnelles) > 0 {
		enc.DonneesAdditionnelles = &additionnelles
	}

	return nil
}

// ExecuteDeclaration exécute la procédure stockée
func (s *EncoursService) ExecuteDeclaration(ctx context.Context, dateDeclaration string) error {
	if dateDeclaration == "" {
		return fmt.Errorf("date de déclaration is required")
	}

	if err := s.repo.ExecuteStoredProcedure(ctx, dateDeclaration); err != nil {
		log.Printf("Error executing stored procedure: %v", err)
		return err
	}

	return nil
}

// ValidateEncours valide les données d'un encours
func (s *EncoursService) ValidateEncours(enc *models.Encours) error {
	if enc.RefIntEng == "" {
		return fmt.Errorf("RefIntEng cannot be empty")
	}
	return nil
}

// CountEncours retourne le nombre total d'encours
func (s *EncoursService) CountEncours(ctx context.Context) (int, error) {
	encours, err := s.repo.GetAll(ctx)
	if err != nil {
		return 0, err
	}
	return len(encours), nil
}
