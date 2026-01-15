package repository

import (
	"database/sql"
	"fmt"

	"GoWebapitest/internal/core/domain/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	query := r.db.QueryRow(`
		SELECT id, email, motdepasse, role
		FROM public.user
		WHERE email = $1
	`, email)
	var u models.User
	err := query.Scan(&u.Id, &u.Email, &u.MotDePasse, &u.Role)
	if err != nil {
		fmt.Println("Error in FindByEmail:", err)
		return nil, err
	}

	return &u, nil
}
