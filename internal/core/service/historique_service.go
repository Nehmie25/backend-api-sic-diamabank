package service

import (
	"GoWebapitest/internal/core/domain/models"
	"GoWebapitest/internal/core/port"
	"fmt"
)

type HistoriqueService struct {
	repo port.HistoriqueRepository
}

func NewHistoriqueService(r port.HistoriqueRepository) *HistoriqueService {
	return &HistoriqueService{repo: r}
}

func (s *HistoriqueService) LogAction(Operation string, UserID string, UserNames string) error {
	log := &models.Historique{
		Operation: Operation,
		UserID:    UserID,
		UserNames: UserNames,
	}

	fmt.Println("Logging action: ", log)

	return s.repo.Save(log)
}
