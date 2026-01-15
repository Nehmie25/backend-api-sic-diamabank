package repository

import (
	"GoWebapitest/internal/core/domain/models"
	"context"
	"database/sql"
	"time"
)

type encoursRepo struct {
	db *sql.DB
}

func NewEncoursRepo(db *sql.DB) *encoursRepo {
	return &encoursRepo{db: db}
}

func (r *encoursRepo) GenerateEncours(date string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, `BEGIN DIAMA.GEN_ENCOURS(:1); END;`, date)
	return err
}

func (r *encoursRepo) FetchAllEncours() ([]models.Encours, error) {
	rows, err := r.db.Query(`SELECT NatDec,RefIntEng,CodDev,DatEch,MntDerEch,MonPai,DatPai,MntHBil,MntRemAnt,MntCRDU,MntCreRat,MntUtilise,TO_CHAR(MntAgi),MntCapImp,
MntTotImp, DatDefaill, MntPro, MntPerte, NbrEchPay, NbrEchImp, NbrEchRest, QualiCre, PD, LGD, CCF, IFRSStage, DatEvent FROM DIAMA.DBANK_ENCOURS`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var result []models.Encours
	var enc models.Encours
	for rows.Next(){
		rows.Scan(&enc.NatDec, &enc.RefIntEng, &enc.CodDev, &enc.DatEch, &enc.MntDerEch, &enc.MonPai, &enc.DatPai, &enc.MntHBil, &enc.MntRemAnt, &enc.MntCRDU,
			&enc.MntCreRat, &enc.MntUtilise, &enc.MntAgi, &enc.MntCapImp, &enc.MntTotImp, &enc.DatDefaill, &enc.MntPro, &enc.MntPerte, &enc.NbrEchPay, &enc.NbrEchImp,
			&enc.NbrEchRest, &enc.QualiCre, &enc.PD, &enc.LGD, &enc.CCF, &enc.IFRSStage, &enc.DatEvent)

		// Données additionnelles
		var additionnelles []models.EncoursAdditionnelles
		var additionnelle models.EncoursAdditionnelles

		rowa, err := r.db.Query("SELECT Cle, Valeur FROM DIAMA.DBANK_ENC_ADDITIONNELLES WHERE RefIntEng = :1", enc.RefIntEng)
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
			enc.DonneesAdditionnelles = &additionnelles
		}
		// end données additionnelles

		result = append(result, enc)
	}

	return result, nil
}
