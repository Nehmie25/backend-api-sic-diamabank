package services

import (
	"context"
	"fmt"
	"log"
	"sicbcrg.diamabank.com/internal/models"
	"sicbcrg.diamabank.com/internal/repositories"
)

// PersonnePhysiqueService contient la logique métier
type PersonnePhysiqueService struct {
	repo repositories.PersonnePhysiqueRepository
}

// NewPersonnePhysiqueService crée une nouvelle instance du service
func NewPersonnePhysiqueService(repo repositories.PersonnePhysiqueRepository) *PersonnePhysiqueService {
	return &PersonnePhysiqueService{
		repo: repo,
	}
}

// GetAllPersonnesWithDetails récupère toutes les personnes avec leurs données associées
func (s *PersonnePhysiqueService) GetAllPersonnesWithDetails(ctx context.Context) ([]models.PersonnePhysique, error) {
	// Récupérer toutes les personnes depuis le repository
	personnes, err := s.repo.GetAll(ctx)
	if err != nil {
		log.Printf("Error fetching personnes: %v", err)
		return nil, err
	}

	// Enrichir chaque personne avec ses relations
	for i := range personnes {
		if err := s.enrichirPersonne(ctx, &personnes[i]); err != nil {
			log.Printf("Error enriching personne %s: %v", personnes[i].IdInterneClt, err)
			// Continuer même en cas d'erreur pour une personne
			continue
		}
	}

	return personnes, nil
}

// enrichirPersonne charge tous les détails associés à une personne
func (s *PersonnePhysiqueService) enrichirPersonne(ctx context.Context, pp *models.PersonnePhysique) error {
	// Charger comptes associés
	comptes, err := s.repo.GetComptes(ctx, pp.IdInterneClt)
	if err != nil {
		log.Printf("Error loading comptes: %v", err)
	} else if len(comptes) > 0 {
		pp.CompteAssocie = &comptes
	}

	// Charger pièces d'identité
	pieces, err := s.repo.GetPieces(ctx, pp.IdInterneClt)
	if err != nil {
		log.Printf("Error loading pieces: %v", err)
	} else if len(pieces) > 0 {
		pp.Piece = &pieces
	}

	// Charger données complémentaires
	complementaire, err := s.repo.GetComplementaire(ctx, pp.IdInterneClt)
	if err != nil {
		log.Printf("Error loading complementaire: %v", err)
	} else if complementaire != nil {
		pp.DonneeComplementaire = complementaire
	}

	// Charger tuteurs/curateurs
	tuteurs, err := s.repo.GetTuteurs(ctx, pp.IdInterneClt, pp.STutelle)
	if err != nil {
		log.Printf("Error loading tuteurs: %v", err)
	} else if len(tuteurs) > 0 {
		pp.TuteurCurateur = &tuteurs
	}

	// Charger employeur
	employeur, err := s.repo.GetEmployeur(ctx, pp.IdInterneClt)
	if err != nil {
		log.Printf("Error loading employeur: %v", err)
	} else if employeur != nil {
		pp.Employeur = employeur
	}

	// Charger données additionnelles
	additionnelles, err := s.repo.GetAdditionnelles(ctx, pp.IdInterneClt)
	if err != nil {
		log.Printf("Error loading additionnelles: %v", err)
	} else if additionnelles != nil {
		pp.DonneesAdditionnelles = additionnelles
	}

	return nil
}

// ExecuteDeclaration exécute la procédure stockée et récupère les données
func (s *PersonnePhysiqueService) ExecuteDeclaration(ctx context.Context, dateDeclaration string) error {
	// Valider la date
	if dateDeclaration == "" {
		return fmt.Errorf("date de déclaration is required")
	}

	// Exécuter la procédure stockée
	if err := s.repo.ExecuteStoredProcedure(ctx, dateDeclaration); err != nil {
		log.Printf("Error executing stored procedure: %v", err)
		return err
	}

	return nil
}

// ValidatePersonnePhysique valide les données d'une personne
func (s *PersonnePhysiqueService) ValidatePersonnePhysique(pp *models.PersonnePhysique) error {
	if pp.IdInterneClt == "" {
		return fmt.Errorf("IdInterneClt cannot be empty")
	}
	if pp.NomNaiClt == "" {
		return fmt.Errorf("NomNaiClt cannot be empty")
	}
	if pp.Client == "" {
		return fmt.Errorf("Client cannot be empty")
	}
	return nil
}

// GetPersonneByID récupère une personne spécifique avec ses détails
func (s *PersonnePhysiqueService) GetPersonneByID(ctx context.Context, idInterneClt string) (*models.PersonnePhysique, error) {
	// Pour l'instant, récupérer tous et filtrer
	// À optimiser avec une vraie requête GetByID dans le repository
	personnes, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	for i := range personnes {
		if personnes[i].IdInterneClt == idInterneClt {
			if err := s.enrichirPersonne(ctx, &personnes[i]); err != nil {
				log.Printf("Error enriching personne: %v", err)
			}
			return &personnes[i], nil
		}
	}

	return nil, fmt.Errorf("personne not found with id: %s", idInterneClt)
}

// CountPersonnes retourne le nombre total de personnes physiques
func (s *PersonnePhysiqueService) CountPersonnes(ctx context.Context) (int, error) {
	personnes, err := s.repo.GetAll(ctx)
	if err != nil {
		return 0, err
	}
	return len(personnes), nil
}
