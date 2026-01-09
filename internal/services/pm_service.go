package services

import (
	"context"
	"fmt"
	"log"
	"sicbcrg.diamabank.com/internal/models"
	"sicbcrg.diamabank.com/internal/repositories"
)

// PersonneMoraleService contient la logique métier
type PersonneMoraleService struct {
	repo repositories.PersonneMoraleRepository
}

// NewPersonneMoraleService crée une nouvelle instance du service
func NewPersonneMoraleService(repo repositories.PersonneMoraleRepository) *PersonneMoraleService {
	return &PersonneMoraleService{
		repo: repo,
	}
}

// GetAllPersonnesWithDetails récupère toutes les personnes morales avec leurs données associées
func (s *PersonneMoraleService) GetAllPersonnesWithDetails(ctx context.Context) ([]models.PersonneMorale, error) {
	personnes, err := s.repo.GetAll(ctx)
	if err != nil {
		log.Printf("Error fetching personnes morales: %v", err)
		return nil, err
	}

	// Enrichir chaque personne morale avec ses relations
	for i := range personnes {
		if err := s.enrichirPersonne(ctx, &personnes[i]); err != nil {
			log.Printf("Error enriching personne morale %s: %v", personnes[i].IdInterneClt, err)
			continue
		}
	}

	return personnes, nil
}

// enrichirPersonne charge tous les détails associés à une personne morale
func (s *PersonneMoraleService) enrichirPersonne(ctx context.Context, pm *models.PersonneMorale) error {
	// Charger mandataires
	mandataires, err := s.repo.GetMandataires(ctx, pm.IdInterneClt)
	if err != nil {
		log.Printf("Error loading mandataires: %v", err)
	} else if len(mandataires) > 0 {
		pm.Mandataires = &mandataires
	}

	// Charger comptes associés
	comptes, err := s.repo.GetComptesAssocies(ctx, pm.IdInterneClt)
	if err != nil {
		log.Printf("Error loading comptes: %v", err)
	} else if len(comptes) > 0 {
		// Enrichir les comptes avec mandataires associés
		for j := range comptes {
			mandatairesAssocie, err := s.repo.GetMandatairesAssocie(ctx, comptes[j].IdInterneClt)
			if err == nil && len(mandatairesAssocie) > 0 {
				comptes[j].MandataireAssocie = &mandatairesAssocie
			}
		}
		pm.CompteAssocie = &comptes
	}

	// Charger actionnaires
	actionnaires, err := s.repo.GetActionnaires(ctx, pm.IdInterneClt)
	if err != nil {
		log.Printf("Error loading actionnaires: %v", err)
	} else if len(actionnaires) > 0 {
		pm.Actionnaire = &actionnaires
	}

	// Charger données additionnelles
	additionnelles, err := s.repo.GetAdditionnelles(ctx, pm.IdInterneClt)
	if err != nil {
		log.Printf("Error loading additionnelles: %v", err)
	} else if len(additionnelles) > 0 {
		pm.DonneesAdditionnelles = &additionnelles
	}

	return nil
}

// ExecuteDeclaration exécute la procédure stockée et récupère les données
func (s *PersonneMoraleService) ExecuteDeclaration(ctx context.Context, dateDeclaration string) error {
	if dateDeclaration == "" {
		return fmt.Errorf("date de déclaration is required")
	}

	if err := s.repo.ExecuteStoredProcedure(ctx, dateDeclaration); err != nil {
		log.Printf("Error executing stored procedure: %v", err)
		return err
	}

	return nil
}

// ValidatePersonneMorale valide les données
func (s *PersonneMoraleService) ValidatePersonneMorale(pm *models.PersonneMorale) error {
	if pm.IdInterneClt == "" {
		return fmt.Errorf("IdInterneClt cannot be empty")
	}
	if pm.DenomSocial == "" {
		return fmt.Errorf("DenomSocial cannot be empty")
	}
	return nil
}

// CountPersonnes retourne le nombre total de personnes morales
func (s *PersonneMoraleService) CountPersonnes(ctx context.Context) (int, error) {
	personnes, err := s.repo.GetAll(ctx)
	if err != nil {
		return 0, err
	}
	return len(personnes), nil
}
