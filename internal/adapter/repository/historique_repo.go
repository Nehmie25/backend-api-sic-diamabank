package repository

import (
	"database/sql"
	"fmt"
	"strconv"

	"GoWebapitest/internal/core/domain/models"
)

type HistoriqueRepository struct {
	db *sql.DB
}

func NewHistoriqueRepository(db *sql.DB) *HistoriqueRepository {
	return &HistoriqueRepository{db: db}
}

func (r *HistoriqueRepository) Save(log *models.Historique) error {
	fmt.Println("Saving log:", log)
	UserId := log.UserID
	UserIdInt, _ := strconv.Atoi(UserId)
	_, err := r.db.Exec(`
		INSERT INTO public.historique
		(operation,userid,usernames) VALUES ($1, $2, $3)
	`,
		log.Operation,
		UserIdInt,
		log.UserNames,
	)
	if err != nil {
		return err
	}
	return nil
}
