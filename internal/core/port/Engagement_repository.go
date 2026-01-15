package port

import "GoWebapitest/internal/core/domain/models"

type EngagementRepository interface {
	GenerateEngagements(date string) error
	FetchAllEngagements() ([]models.Engagements, error)
}
