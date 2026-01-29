package repository

import (
	"database/sql"
)

type ResetPasswordRepo struct {
	db *sql.DB
}

func NewResetPasswordRepo(db *sql.DB) *ResetPasswordRepo {
	return &ResetPasswordRepo{db: db}
}

func (r *ResetPasswordRepo) ResetPassword(userid int, Password string) error {
	_,err := r.db.Exec(`
		UPDATE public.user
		SET motdepasse = $1
		WHERE id = $2
	`, Password, userid)
	if err != nil {
		return err
	}

	return nil
}