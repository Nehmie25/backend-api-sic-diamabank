package port

import "GoWebapitest/internal/core/domain/models"

type HistoriqueRepository interface {
	Save(log *models.Historique) error
	FindAllLogs(offset int) ([]*models.Historique, error)
	Count() (int, error)
}
