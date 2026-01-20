package port

import "GoWebapitest/internal/core/domain/models"

type UserRepository interface {
	FindByEmail(email string) (*models.User, error)
	FindAll() ([]models.User, error)
	AddUser(nom, email, motdepasse, role string) error
	ModifyStatus(id string, isactive bool) error
}
