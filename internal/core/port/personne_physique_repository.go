package port

import "GoWebapitest/internal/core/domain/models"

type PersonnePhysiqueRepository interface {
	GeneratePersonnesPhysique(date string) error
	FetchAllPersonnesPhysiques() ([]models.PersonnePhysique, error)
}
