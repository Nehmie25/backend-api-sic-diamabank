package repository

import (
	"GoWebapitest/internal/core/domain/models"
	"context"
	"database/sql"
	"time"
)

type PersonneMoraleRepo struct {
	db *sql.DB
}

func NewPersonneMoralRepo(db *sql.DB) *PersonneMoraleRepo {
	return &PersonneMoraleRepo{db: db}
}

func (r *PersonneMoraleRepo) GeneratePersonnesMoral(date string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, `BEGIN DIAMA.GEN_PERSMORALE(:1); END;`, date)
	return err
}

func (r *PersonneMoraleRepo) FetchAllPersonnesMoral() ([]models.PersonneMorale, error) {
	rows, err := r.db.Query("SELECT * FROM DIAMA.DBANK_PERSMORALE")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var result []models.PersonneMorale
	var pm models.PersonneMorale
	for rows.Next() {
		rows.Scan(&pm.Client, &pm.EstDeclare, &pm.DateDeclare, &pm.Datmaj, &pm.NatDec, &pm.NatClient, &pm.IdInterneClt, &pm.DenomSocial,
			&pm.Sigle, &pm.DatCreat, &pm.Statut, &pm.DatCreaPart, &pm.FormeJuridique, &pm.PaysSiegeSocial, &pm.VilleSiegeSocial, &pm.Mobile, &pm.Email,
			&pm.SiteWeb, &pm.Adress, &pm.CommuneAdress, &pm.CodePostal, &pm.Resident, &pm.RCCM, &pm.NIF, &pm.NIFP, &pm.NumAgrement, &pm.NumSecSoc,
			&pm.SecActEcon, &pm.SectInst, &pm.SitBancaire, &pm.DateDebIB, &pm.DateFinIB)

		//cas Mandataires
		rowc, err := r.db.Query("SELECT * FROM DIAMA.DBANK_MANDATAIRE WHERE IdInterneClt = :1", pm.IdInterneClt)
		if err != nil {
			return nil, err
		}

		defer rowc.Close()

		var mandataires []models.PersonneMoraleMandataire
		for rowc.Next() {
			var mandataire models.PersonneMoraleMandataire
			if err := rowc.Scan(&mandataire.Client, &mandataire.Idp, &mandataire.EstDeclare, &mandataire.DateDeclare, &mandataire.Datmaj, &mandataire.IdInterneClt,
				&mandataire.IdInterneMdt, &mandataire.Qualite, &mandataire.DatDebMdt, &mandataire.DatFinMdt); err != nil {
				return nil, err
			}
			mandataires = append(mandataires, mandataire)
		}
		if len(mandataires) > 0 {
			pm.Mandataires = &mandataires
		}
		//end mandataires

		//cas des comptes associes
		rowp, err := r.db.Query("SELECT * FROM DIAMA.DBANK_PERSMORALE_CPT WHERE IdInterneClt = :1", pm.IdInterneClt)
		if err != nil {
			return nil, err
		}
		defer rowp.Close()
		var comptesassocies []models.PersonneMoraleCompteAssocie
		for rowp.Next() {
			var compteassocie models.PersonneMoraleCompteAssocie
			if err := rowp.Scan(&compteassocie.Client, &compteassocie.Compte, &compteassocie.Devise, &compteassocie.Ncg, &compteassocie.EstDeclare,
				&compteassocie.DateDeclare, &compteassocie.Datmaj, &compteassocie.IdInterneClt, &compteassocie.CodAgce, &compteassocie.NumCpt,
				&compteassocie.CleRib, &compteassocie.StatCpt); err != nil {
				return nil, err
			}

			compteassocie.MandataireAssocie = nil
			//cas mandataires associes au compte
			rowm, err := r.db.Query("SELECT * FROM DIAMA.DBANK_MANDASSOCIE WHERE IdInterneClt = :1", compteassocie.IdInterneClt)
			if err != nil {
				return nil, err
			}
			defer rowm.Close()
			var mandatairesassocies []models.PersonneMoraleMandataireMandataireAssocie
			for rowm.Next() {
				var mandataireassocie models.PersonneMoraleMandataireMandataireAssocie
				if err := rowm.Scan(&mandataireassocie.Client, &mandataireassocie.Idp, &mandataireassocie.EstDeclare, &mandataireassocie.DateDeclare,
					&mandataireassocie.Datmaj, &mandataireassocie.IdInterneClt, &mandataireassocie.IdInterneMdtCpt, &mandataireassocie.DatDebMdt,
					&mandataireassocie.DatFinMdt); err != nil {
					return nil, err
				}
				mandatairesassocies = append(mandatairesassocies, mandataireassocie)
			}
			if len(mandatairesassocies) > 0 {
				compteassocie.MandataireAssocie = &mandatairesassocies
			}

			//end mandataires associes au compte
			comptesassocies = append(comptesassocies, compteassocie)
		}
		if len(comptesassocies) > 0 {
			pm.CompteAssocie = &comptesassocies
		}
		//end comptes associes

		//cas actionnaires
		rowa, err := r.db.Query("SELECT * FROM DIAMA.DBANK_ACTIONNAIRE WHERE IdInterneClt = :1", pm.IdInterneClt)
		if err != nil {
			return nil, err
		}
		defer rowa.Close()
		var actionnaires []models.PersonneMoraleActionnaire
		for rowa.Next() {
			var actionnaire models.PersonneMoraleActionnaire
			if err := rowa.Scan(&actionnaire.Client, &actionnaire.Idp, &actionnaire.EstDeclare, &actionnaire.DateDeclare, &actionnaire.Datmaj,
				&actionnaire.IdInterneClt, &actionnaire.IdInterneAct, &actionnaire.PartAct, &actionnaire.DaEntrAct); err != nil {
				return nil, err
			}
			actionnaires = append(actionnaires, actionnaire)
		}

		//end actionnaires
		if len(actionnaires) > 0 {
			pm.Actionnaire = &actionnaires
		}

		//cas des donnees additionnelles
		rowd, err := r.db.Query("SELECT * FROM DIAMA.DBANK_PM_ADDITIONNELLES WHERE IdInterneClt = :1", pm.IdInterneClt)
		if err != nil {
			return nil, err
		}
		defer rowd.Close()
		var additionnelles []models.PersonneMoraleAdditionnelles
		for rowd.Next() {
			var additionnelle models.PersonneMoraleAdditionnelles
			if err := rowd.Scan(&additionnelle.Client, &additionnelle.EstDeclare, &additionnelle.DateDeclare, &additionnelle.Datmaj,
				&additionnelle.IdInterneClt, &additionnelle.Cle, &additionnelle.Valeur); err != nil {
				return nil, err
			}
			additionnelles = append(additionnelles, additionnelle)
		}
		if len(additionnelles) > 0 {
			pm.DonneesAdditionnelles = &additionnelles
		}
		//end donnees additionnelles
		result = append(result, pm)
	}

	return result, nil
}
