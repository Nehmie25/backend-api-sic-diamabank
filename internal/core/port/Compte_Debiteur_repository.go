package port

import "GoWebapitest/internal/core/domain/models"

type CompteDebiteursRepository interface {
	GenerateEncours(date string) error
	FetchAllEncours() ([]models.CompteDebiteurs, error)
}
