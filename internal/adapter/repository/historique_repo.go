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

func (r *HistoriqueRepository) FindAllLogs(offset int) ([]*models.Historique, error) {
	rows,err := r.db.Query(`SELECT * FROM public.historique LIMIT 50 OFFSET $1`, offset)
	if err != nil{
		return nil, err
	}
	defer rows.Close()
	var logs []*models.Historique
	for rows.Next(){
		var log models.Historique
		err := rows.Scan(&log.ID, &log.Operation, &log.UserID, &log.UserNames, &log.Date, &log.Time)
		if err != nil{
			fmt.Println("Error scanning log:", err)
			return nil, err
		}
		logs = append(logs, &log)
	}
	return logs, nil
}

func (r *HistoriqueRepository) Count() (int, error) {
	var total int
	err := r.db.QueryRow(`SELECT COUNT(*) FROM historique`).Scan(&total)
	return total, err
}

