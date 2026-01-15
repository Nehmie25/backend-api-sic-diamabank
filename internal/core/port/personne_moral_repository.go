package port

import "GoWebapitest/internal/core/domain/models"

type PersonneMoralRepository interface {
	GeneratePersonnesMoral(date string) error
	FetchAllPersonnesMoral() ([]models.PersonneMorale, error)
}
