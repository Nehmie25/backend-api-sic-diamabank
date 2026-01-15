package repository

import (
	"GoWebapitest/internal/core/domain/models"
	"context"
	"database/sql"
	"strconv"
	"time"
)

type engagementsRepo struct {
	db *sql.DB
}

func NewEngagementRepo(db *sql.DB) *engagementsRepo {
	return &engagementsRepo{db: db}
}

func (r *engagementsRepo) GenerateEngagements(date string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, `BEGIN DIAMA.GEN_ENGAGEMENTS(:1); END;`, date)
	return err
}

func (r *engagementsRepo) FetchAllEngagements() ([]models.Engagements, error) {
	rows, err := r.db.Query(`SELECT RefIntEng, NatDec, TypEve, LigneParent, RefIntLigne, RefDemandeEng, DatDem, TypModif, EstDout, 
			Cloture, MotifCloture, DatClo,DatAccord, DateMEP, TypEng, TO_CHAR(MntEng), TO_CHAR(MntInt), CodDev, 
			PeriodRemb, TxIntEng, TypTxInt, TxComm, IndRef, Sprd, TxEffGlob, MoyRemb, TypAmo, TypDiffAmo, UnitDur, PerDiffAmo,
			MntEch, NbrEch, DatPremEch, DatFin, MntFrais, MntComm, CodAgce, EstRachatCreance,
			ParCont, TO_CHAR(ValNom), TO_CHAR(ValCess), DatEvent, IdIntBen, PourBenef FROM DIAMA.DBANK_ENGAGEMENTS`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var result []models.Engagements
	var eng models.Engagements
	for rows.Next(){
		rows.Scan(&eng.RefIntEng, &eng.NatDec, &eng.TypEve,
			&eng.LigneParent, &eng.RefIntLigne, &eng.RefDemandeEng, &eng.DatDem, &eng.TypeModif, &eng.EstDout, &eng.Cloture, &eng.DatAccord, &eng.MotifCloture, &eng.DatClo,
			&eng.DateMEP, &eng.TypEng, &eng.MntEng, &eng.MntInt, &eng.CodDev, &eng.PeriodRemb, &eng.TxIntEng, &eng.TypTxInt,
			&eng.TxComm, &eng.IndRef, &eng.Sprd, &eng.TxEffGlob, &eng.MoyRemb, &eng.TypAmo, &eng.TypDiffAmo, &eng.UnitDur, &eng.PerDiffAmo,
			&eng.MntEch, &eng.NbrEch, &eng.DatPremEch, &eng.DatFin, &eng.MntFrais, &eng.MntComm, &eng.CodAgce, &eng.EstRachatCreance,
			&eng.ParCont, &eng.ValNom, &eng.ValCess, &eng.DatEvent, &eng.IdIntBen, &eng.PourBenef)


		//end beneficiaire
		var beneficiaire models.EngagementsBeneficiaire
		var beneficiaires []models.EngagementsBeneficiaire = nil

		beneficiaire.IdIntBen = eng.IdIntBen
		beneficiaire.PourBenef, _ = strconv.ParseFloat(eng.PourBenef, 64)
		beneficiaires = append(beneficiaires, beneficiaire)

		if len(beneficiaires) > 0 {
			eng.Beneficiaire = &beneficiaires
		}
		//end beneficiaire

		// Données additionnelles
		var additionnelle models.EngagementsAdditionnelles
		var additionnelles []models.EngagementsAdditionnelles

		rowa, err := r.db.Query("SELECT Cle, Valeur FROM DIAMA.DBANK_ENG_ADDITIONNELLES WHERE RefIntEng = :1", eng.RefIntEng)
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
			eng.DonneesAdditionnelles = &additionnelles
		}
		// end données additionnelles

		result = append(result, eng)
	}

	return result, nil
}
