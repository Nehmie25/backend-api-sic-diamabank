package service

import (
	"GoWebapitest/internal/core/domain/models"
	"GoWebapitest/internal/core/port"
	"math"
)

type HistoriqueService struct {
	repo port.HistoriqueRepository
}

func NewHistoriqueService(r port.HistoriqueRepository) *HistoriqueService {
	return &HistoriqueService{repo: r}
}

func (s *HistoriqueService) LogAction(Operation string, UserID string, UserNames string) error {
	switch Operation { 
	case "/users/historique":
		Operation = "Consultation des historiques"
	case "/users/login":
		Operation = "Connexion utilisateur"
	case "/users/updatestate":
		Operation = "Modification état utilisateur"
	case "/users/register":
		Operation = "Ajout nouvel utilisateur"
	case "/users/":
		Operation = "Consultation des utilisateurs"
	case "/declarations/personnephysique":
		Operation = "Consultation des déclarations personnes physiques"
	case "/declarations/personnemorale":
		Operation = "Consultation des déclarations personnes morales"
	case "/declarations/engagements":
		Operation = "Consultation des déclarations engagements"
	case "/declarations/encours":
		Operation = "Consultation des déclarations encours"
	case "/users/restpassword":
		Operation = "Réinitialisation mot de passe utilisateur"
	case "/declarations/comptedebiteurs":
		Operation = "Consultation des déclarations compte débiteurs"
	default:
		Operation = "Opération inconnue"
	}
	log := &models.Historique{
		Operation: Operation,
		UserID:    UserID,
		UserNames: UserNames,
	}

	return s.repo.Save(log)
}

func (s *HistoriqueService) GetAllLogs(page int) ([]*models.Historique, *models.Pagination, error) {
	limit := 50

	offset := (page - 1) * limit

	total, err := s.repo.Count()
	if err != nil {
		return nil, nil, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	historiques, err := s.repo.FindAllLogs(offset)
	if err != nil {
		return nil, nil, err
	}
		pagination := &models.Pagination{
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: totalPages,
	}
	return historiques, pagination, nil
}