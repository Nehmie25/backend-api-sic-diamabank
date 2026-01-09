package services

import (
	"context"
	"fmt"
	"log"
	"sicbcrg.diamabank.com/internal/models"
	"sicbcrg.diamabank.com/internal/repositories"
)

// CompteDebiteurService contient la logique métier
type CompteDebiteurService struct {
	repo repositories.CompteDebiteurRepository
}

// NewCompteDebiteurService crée une nouvelle instance du service
func NewCompteDebiteurService(repo repositories.CompteDebiteurRepository) *CompteDebiteurService {
	return &CompteDebiteurService{
		repo: repo,
	}
}

// GetAllComptesWithDetails récupère tous les comptes débiteurs avec leurs données associées
func (s *CompteDebiteurService) GetAllComptesWithDetails(ctx context.Context) ([]models.CompteDebiteurs, error) {
	comptes, err := s.repo.GetAll(ctx)
	if err != nil {
		log.Printf("Error fetching comptes debiteurs: %v", err)
		return nil, err
	}

	// Enrichir chaque compte avec ses relations
	for i := range comptes {
		if err := s.enrichirCompte(ctx, &comptes[i]); err != nil {
			log.Printf("Error enriching compte %s: %v", comptes[i].IdIntTit, err)
			continue
		}
	}

	return comptes, nil
}

// enrichirCompte charge tous les détails associés à un compte débiteur
func (s *CompteDebiteurService) enrichirCompte(ctx context.Context, deb *models.CompteDebiteurs) error {
	// Charger le titulaire
	titulaire := models.CompteDebiteursTitulaire{
		IdIntTit: deb.IdIntTit,
	}
	deb.Titulaire = titulaire

	// Charger données additionnelles
	additionnelles, err := s.repo.GetAdditionnelles(ctx, deb.IdIntTit)
	if err != nil {
		log.Printf("Error loading additionnelles: %v", err)
	} else if len(additionnelles) > 0 {
		deb.DonneesAdditionnelles = &additionnelles
	}

	return nil
}

// ExecuteDeclaration exécute la procédure stockée
func (s *CompteDebiteurService) ExecuteDeclaration(ctx context.Context, dateDeclaration string) error {
	if dateDeclaration == "" {
		return fmt.Errorf("date de déclaration is required")
	}

	if err := s.repo.ExecuteStoredProcedure(ctx, dateDeclaration); err != nil {
		log.Printf("Error executing stored procedure: %v", err)
		return err
	}

	return nil
}

// ValidateCompteDebiteur valide les données d'un compte
func (s *CompteDebiteurService) ValidateCompteDebiteur(deb *models.CompteDebiteurs) error {
	if deb.IdIntTit == "" {
		return fmt.Errorf("IdIntTit cannot be empty")
	}
	if deb.Rib == "" {
		return fmt.Errorf("Rib cannot be empty")
	}
	return nil
}

// CountComptes retourne le nombre total de comptes débiteurs
func (s *CompteDebiteurService) CountComptes(ctx context.Context) (int, error) {
	comptes, err := s.repo.GetAll(ctx)
	if err != nil {
		return 0, err
	}
	return len(comptes), nil
}
