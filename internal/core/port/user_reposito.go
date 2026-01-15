package port

import "GoWebapitest/internal/core/domain/models"

type UserRepository interface {
	FindByEmail(email string) (*models.User, error)
}
