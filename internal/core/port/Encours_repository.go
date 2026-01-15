package port

import "GoWebapitest/internal/core/domain/models"

type EncoursRepository interface {
	GenerateEncours(date string) error
	FetchAllEncours() ([]models.Encours, error)
}
