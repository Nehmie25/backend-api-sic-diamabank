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
func (r *UserRepository) FindAll() ([]models.User, error) {
	rows, err := r.db.Query(`
		SELECT id, nom, email, role
		FROM public.user
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		err := rows.Scan(&u.Id, &u.Nom, &u.Email, &u.Role)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

func (r *UserRepository) AddUser(nom, email, motdepasse, role string) (error) {
	_, err := r.db.Exec(`
		INSERT INTO public.user (nom,email,motdepasse,role,isactive) 
		VALUES ($1,$2,$3,$4,$5)
	`, nom, email, motdepasse, role, true)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) ModifyStatus(id int, isactive bool) (error) {
	_, err := r.db.Exec(`
		UPDATE public.user SET isactive=$1
		WHERE id = $2
	`, isactive, id)
	if err != nil {
		return err
	}
	return nil
}