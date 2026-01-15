package repository

import (
	"GoWebapitest/internal/core/domain/models"
	"context"
	"database/sql"
	"time"
)

type CompteDebiteursRepoRepo struct {
	db *sql.DB
}

func NewCompteDebiteursRepo(db *sql.DB) *CompteDebiteursRepoRepo {
	return &CompteDebiteursRepoRepo{db: db}
}

func (r *CompteDebiteursRepoRepo) GenerateEncours(date string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, `BEGIN DIAMA.GEN_CPTDEBITEURS(:1); END;`, date)
	return err
}

func (r *CompteDebiteursRepoRepo) FetchAllEncours() ([]models.CompteDebiteurs, error) {
	rows, err := r.db.Query(`SELECT NatDec,CodDev,Rib,TO_CHAR(SoldeDeb),DateDefaill,NbrJourDebMax,TO_CHAR(SoldeDebMax),TO_CHAR(MntProv),TO_CHAR(MntPerte),
			TO_CHAR(MntAgi),TO_CHAR(QualiCre),IdIntTit	FROM DIAMA.DBANK_CPTDEBITEURS`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var result []models.CompteDebiteurs
	var deb models.CompteDebiteurs
	for rows.Next(){
		rows.Scan(&deb.NatDec, &deb.CodDev, &deb.Rib, &deb.SoldeDeb, &deb.DateDefaill, &deb.NbrJourDebMax, &deb.SoldeDebMax,
			&deb.MntProv, &deb.MntPerte, &deb.MntAgi, &deb.QualiCre, &deb.IdIntTit)

		var titulaire models.CompteDebiteursTitulaire
		titulaire.IdIntTit = deb.IdIntTit
		deb.Titulaire = titulaire

		var additionnelles []models.CompteDebiteursAdditionnelles
		var additionnelle models.CompteDebiteursAdditionnelles

		// Données additionnelles du titulaire
		rowa, err := r.db.Query("SELECT Cle, Valeur FROM DIAMA.DBANK_DEB_ADDITIONNELLES WHERE RefIntEng = :1", deb.IdIntTit)
		if err != nil {
			return nil, err
		}
		defer rowa.Close()

		for rowa.Next() {
			if err := rowa.Scan(&additionnelle.Cle, &additionnelle.Valeur); err != nil {
				return nil, err
			}
			additionnelles = append(additionnelles, additionnelle)
		}
		if len(additionnelles) > 0 {
			deb.DonneesAdditionnelles = &additionnelles
		}
		// end données additionnelles

		result = append(result, deb)
	}

	return result, nil
}
