package repository

import (
	"context"
	"database/sql"
	"time"
	"GoWebapitest/internal/core/domain/models"
)

type PersonnePhysiqueRepo struct {
	db *sql.DB
}

func NewPersonnePhysiqueRepo(db *sql.DB) *PersonnePhysiqueRepo {
	return &PersonnePhysiqueRepo{db: db}
}

func (r *PersonnePhysiqueRepo) GeneratePersonnesPhysique(date string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, `BEGIN DIAMA.GEN_PERSPHYSIQUE(:1); END;`, date)
	return err
}

func (r *PersonnePhysiqueRepo) FetchAllPersonnesPhysiques() ([]models.PersonnePhysique, error) {
	rows, err := r.db.Query(`SELECT Client,Idp,EstDeclare,DateDeclare,Datmaj,NatDec,NatClient,NIN,IdInterneClt,DatCreaPart,NomNaiClt,NomMtlClt,PrenomClt,Sexe,DatNai,EtatCivil,NomPere,PrenomPere,NomNaiMere,PrmMre,VilleNai,PaysNai,NatClt,Resident,PaysRes,Mobile,Email,Adress,CommuneAdress,CodePostal,Profession,SecActEcon,SectInst,STutelle,StatutClt,DateDeces,SitBancaire,DateDebIB,DateFinIB FROM DIAMA.DBANK_PERSPHYSIQUE`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.PersonnePhysique
	var pp models.PersonnePhysique
	for rows.Next() {

		rows.Scan(&pp.Client, &pp.Idp, &pp.EstDeclare, &pp.DateDeclare, &pp.Datmaj,
			&pp.NatDec, &pp.NatClient, &pp.NIN, &pp.IdInterneClt, &pp.DatCreaPart, &pp.NomNaiClt, &pp.NomMtlClt,
			&pp.PrenomClt, &pp.Sexe, &pp.DatNai, &pp.EtatCivil, &pp.NomPere, &pp.PrenomPere, &pp.NomNaiMere,
			&pp.PrmMre, &pp.VilleNai, &pp.PaysNai, &pp.NatClt, &pp.Resident, &pp.PaysRes, &pp.Mobile,
			&pp.Email, &pp.Adress, &pp.CommuneAdress, &pp.CodePostal, &pp.Profession, &pp.SecActEcon, &pp.SectInst,
			&pp.STutelle, &pp.StatutClt, &pp.DateDeces, &pp.SitBancaire, &pp.DateDebIB, &pp.DateFinIB)
		

		//cas des comptes associes
		rowc, err := r.db.Query("SELECT CodAgce,NumCpt,CleRib,TypCpt,StatCpt FROM DIAMA.DBANK_PERSPHYSIQUE_CPT WHERE IdInterneClt = :1", pp.IdInterneClt)
		if err != nil {
			return nil,err
		}

		defer rowc.Close()

		var comptes []models.PersonnePhysiqueCompteAssocie
		for rowc.Next() {
			var compte models.PersonnePhysiqueCompteAssocie
			if err := rowc.Scan(&compte.CodAgce, &compte.NumCpt, &compte.CleRib, &compte.TypCpt, &compte.StatCpt); err != nil {
				return nil,err
			}
			comptes = append(comptes, compte)
		}
		if len(comptes) > 0 {
			pp.CompteAssocie = &comptes
		}
		//end comptes associes

		//cas des pieces
		rowp, err := r.db.Query("SELECT TypPiece,NumPiece,DatEmiPiece,LieuEmiPiece,PaysEmiPiece,FinValPiece FROM DIAMA.DBANK_PP_PIECE WHERE IdInterneClt = :1", pp.IdInterneClt)
		if err != nil {
			return nil,err
		}

		defer rowp.Close()

		var pieces []models.PersonnePhysiquePiece

		for rowp.Next() {
			var piece models.PersonnePhysiquePiece
			if err := rowp.Scan(&piece.TypPiece, &piece.NumPiece, &piece.DatEmiPiece, &piece.LieuEmiPiece, &piece.PaysEmiPiece, &piece.FinValPiece); err != nil {
				return nil,err
			}
			pieces = append(pieces, piece)
		}
		if len(pieces) > 0 {
			pp.Piece = &pieces
		}
		//end pieces

		//cas des donnees complementaires
		rowd, err := r.db.Query("SELECT NbPersCharge,RevMensMoy,DepMensMoy,PropLoc FROM DIAMA.DBANK_PP_COMPLEMENTAIRE WHERE IdInterneClt = :1", pp.IdInterneClt)
		if err != nil {
			return nil,err
		}
		defer rowd.Close()
		var donnee models.PersonnePhysiqueComplementaire
		if rowd.Next() {
			if err := rowd.Scan(&donnee.NbPersCharge, &donnee.RevMensMoy, &donnee.DepMensMoy, &donnee.PropLoc); err != nil {
				return nil,err
			}
			pp.DonneeComplementaire = &donnee
		}
		//end donnees complementaires

		//cas des tuteurs curateurs
		if pp.STutelle != "O" {
			rowt, err := r.db.Query("SELECT IdInterneMdt,Qualite,DatDbtMdt FROM DIAMA.DBANK_TUTEURCURATEUR WHERE IdInterneClt = :1", pp.IdInterneClt)
			if err != nil {
				return nil,err
			}

			defer rowt.Close()

			var tuteurs []models.PersonnePhysiqueTuteurCurateur

			for rowt.Next() {
				var tuteur models.PersonnePhysiqueTuteurCurateur
				if err := rowt.Scan(&tuteur.IdInterneMdt, &tuteur.Qualite, &tuteur.DatDbtMdt); err != nil {
					return nil,err
				}
				if tuteur.IdInterneMdt != "" {
					tuteurs = append(tuteurs, tuteur)
				}
			}
			if len(tuteurs) > 0 {
				pp.TuteurCurateur = &tuteurs
			}
		}
		//end tuteurs curateurs

		//cas des employeurs
		rowe, err := r.db.Query("SELECT IdInterneEmpl,DenominationSociale,RCCM,NIF,NIFP,DateCreation,DateEntree  FROM DIAMA.DBANK_EMPLOYEUR WHERE IdInterneClt = :1", pp.IdInterneClt)
		if err != nil {
			return nil,err
		}
		defer rowe.Close()
		var employeur models.PersonnePhysiqueEmployeur
		if rowe.Next() {
			if err := rowe.Scan(&employeur.IdInterneEmpl, &employeur.DenominationSociale, &employeur.RCCM, &employeur.NIF, &employeur.NIFP, &employeur.DateCreation, &employeur.DateEntree); err != nil {
				return nil,err
			}
			pp.Employeur = &employeur
		}
		//end employeurs

		//cas des donnees additionnelles
		rowa, err := r.db.Query("SELECT Cle,Valeur FROM DIAMA.DBANK_PP_ADDITIONNELLES WHERE IdInterneClt = :1", pp.IdInterneClt)
		if err != nil {
			return nil,err
		}
		defer rowa.Close()
		var additionnelles models.PersonnePhysiqueAdditionnelles
		if rowa.Next() {
			if err := rowa.Scan(&additionnelles.Cle, &additionnelles.Valeur); err != nil {
				return nil,err
			}
			pp.DonneesAdditionnelles = &additionnelles
		}
		//end donnees additionnelles
		result = append(result, pp)
	}

	return result, nil
}
