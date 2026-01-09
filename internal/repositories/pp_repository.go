package repositories

import (
	"context"
	"database/sql"
	"log"
	"sicbcrg.diamabank.com/internal/models"
)

// PersonnePhysiqueRepository définit les opérations de données pour PersonnePhysique
type PersonnePhysiqueRepository interface {
	GetAll(ctx context.Context) ([]models.PersonnePhysique, error)
	GetComptes(ctx context.Context, idInterneClt string) ([]models.PersonnePhysiqueCompteAssocie, error)
	GetPieces(ctx context.Context, idInterneClt string) ([]models.PersonnePhysiquePiece, error)
	GetComplementaire(ctx context.Context, idInterneClt string) (*models.PersonnePhysiqueComplementaire, error)
	GetTuteurs(ctx context.Context, idInterneClt string, sTutelle string) ([]models.PersonnePhysiqueTuteurCurateur, error)
	GetEmployeur(ctx context.Context, idInterneClt string) (*models.PersonnePhysiqueEmployeur, error)
	GetAdditionnelles(ctx context.Context, idInterneClt string) (*models.PersonnePhysiqueAdditionnelles, error)
	ExecuteStoredProcedure(ctx context.Context, dateDeclaration string) error
}

// OraclePersonnePhysiqueRepository implémente PersonnePhysiqueRepository pour Oracle
type OraclePersonnePhysiqueRepository struct {
	db *sql.DB
}

// NewOraclePersonnePhysiqueRepository crée une nouvelle instance
func NewOraclePersonnePhysiqueRepository(db *sql.DB) PersonnePhysiqueRepository {
	return &OraclePersonnePhysiqueRepository{db: db}
}

// GetAll récupère toutes les personnes physiques
func (r *OraclePersonnePhysiqueRepository) GetAll(ctx context.Context) ([]models.PersonnePhysique, error) {
	var personnes []models.PersonnePhysique

	query := `SELECT Client,Idp,EstDeclare,DateDeclare,Datmaj,NatDec,NatClient,NIN,IdInterneClt,
	          DatCreaPart,NomNaiClt,NomMtlClt,PrenomClt,Sexe,DatNai,EtatCivil,NomPere,PrenomPere,
	          NomNaiMere,PrmMre,VilleNai,PaysNai,NatClt,Resident,PaysRes,Mobile,Email,Adress,
	          CommuneAdress,CodePostal,Profession,SecActEcon,SectInst,STutelle,StatutClt,DateDeces,
	          SitBancaire,DateDebIB,DateFinIB 
	          FROM DIAMA.DBANK_PERSPHYSIQUE`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		log.Printf("Error querying PersonnePhysique: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var pp models.PersonnePhysique
		if err := rows.Scan(&pp.Client, &pp.Idp, &pp.EstDeclare, &pp.DateDeclare, &pp.Datmaj,
			&pp.NatDec, &pp.NatClient, &pp.NIN, &pp.IdInterneClt, &pp.DatCreaPart, &pp.NomNaiClt, 
			&pp.NomMtlClt, &pp.PrenomClt, &pp.Sexe, &pp.DatNai, &pp.EtatCivil, &pp.NomPere, 
			&pp.PrenomPere, &pp.NomNaiMere, &pp.PrmMre, &pp.VilleNai, &pp.PaysNai, &pp.NatClt, 
			&pp.Resident, &pp.PaysRes, &pp.Mobile, &pp.Email, &pp.Adress, &pp.CommuneAdress, 
			&pp.CodePostal, &pp.Profession, &pp.SecActEcon, &pp.SectInst, &pp.STutelle, &pp.StatutClt, 
			&pp.DateDeces, &pp.SitBancaire, &pp.DateDebIB, &pp.DateFinIB); err != nil {
			log.Printf("Error scanning PersonnePhysique row: %v", err)
			return nil, err
		}
		personnes = append(personnes, pp)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating PersonnePhysique rows: %v", err)
		return nil, err
	}

	return personnes, nil
}

// GetComptes récupère les comptes associés à une personne physique
func (r *OraclePersonnePhysiqueRepository) GetComptes(ctx context.Context, idInterneClt string) ([]models.PersonnePhysiqueCompteAssocie, error) {
	var comptes []models.PersonnePhysiqueCompteAssocie

	query := `SELECT CodAgce,NumCpt,CleRib,TypCpt,StatCpt 
	          FROM DIAMA.DBANK_PERSPHYSIQUE_CPT 
	          WHERE IdInterneClt = :1`

	rows, err := r.db.QueryContext(ctx, query, idInterneClt)
	if err != nil {
		log.Printf("Error querying comptes for %s: %v", idInterneClt, err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var compte models.PersonnePhysiqueCompteAssocie
		if err := rows.Scan(&compte.CodAgce, &compte.NumCpt, &compte.CleRib, &compte.TypCpt, &compte.StatCpt); err != nil {
			log.Printf("Error scanning compte row: %v", err)
			return nil, err
		}
		comptes = append(comptes, compte)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating comptes rows: %v", err)
		return nil, err
	}

	return comptes, nil
}

// GetPieces récupère les pièces d'une personne physique
func (r *OraclePersonnePhysiqueRepository) GetPieces(ctx context.Context, idInterneClt string) ([]models.PersonnePhysiquePiece, error) {
	var pieces []models.PersonnePhysiquePiece

	query := `SELECT TypPiece,NumPiece,DatEmiPiece,LieuEmiPiece,PaysEmiPiece,FinValPiece 
	          FROM DIAMA.DBANK_PP_PIECE 
	          WHERE IdInterneClt = :1`

	rows, err := r.db.QueryContext(ctx, query, idInterneClt)
	if err != nil {
		log.Printf("Error querying pieces for %s: %v", idInterneClt, err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var piece models.PersonnePhysiquePiece
		if err := rows.Scan(&piece.TypPiece, &piece.NumPiece, &piece.DatEmiPiece, &piece.LieuEmiPiece, 
			&piece.PaysEmiPiece, &piece.FinValPiece); err != nil {
			log.Printf("Error scanning piece row: %v", err)
			return nil, err
		}
		pieces = append(pieces, piece)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating pieces rows: %v", err)
		return nil, err
	}

	return pieces, nil
}

// GetComplementaire récupère les données complémentaires
func (r *OraclePersonnePhysiqueRepository) GetComplementaire(ctx context.Context, idInterneClt string) (*models.PersonnePhysiqueComplementaire, error) {
	query := `SELECT NbPersCharge,RevMensMoy,DepMensMoy,PropLoc 
	          FROM DIAMA.DBANK_PP_COMPLEMENTAIRE 
	          WHERE IdInterneClt = :1`

	row := r.db.QueryRowContext(ctx, query, idInterneClt)
	
	var donnee models.PersonnePhysiqueComplementaire
	if err := row.Scan(&donnee.NbPersCharge, &donnee.RevMensMoy, &donnee.DepMensMoy, &donnee.PropLoc); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Printf("Error scanning complementaire for %s: %v", idInterneClt, err)
		return nil, err
	}

	return &donnee, nil
}

// GetTuteurs récupère les tuteurs/curateurs
func (r *OraclePersonnePhysiqueRepository) GetTuteurs(ctx context.Context, idInterneClt string, sTutelle string) ([]models.PersonnePhysiqueTuteurCurateur, error) {
	var tuteurs []models.PersonnePhysiqueTuteurCurateur

	// Ne récupérer les tuteurs que si STutelle != "O"
	if sTutelle == "O" {
		return tuteurs, nil
	}

	query := `SELECT IdInterneMdt,Qualite,DatDbtMdt 
	          FROM DIAMA.DBANK_TUTEURCURATEUR 
	          WHERE IdInterneClt = :1`

	rows, err := r.db.QueryContext(ctx, query, idInterneClt)
	if err != nil {
		log.Printf("Error querying tuteurs for %s: %v", idInterneClt, err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tuteur models.PersonnePhysiqueTuteurCurateur
		if err := rows.Scan(&tuteur.IdInterneMdt, &tuteur.Qualite, &tuteur.DatDbtMdt); err != nil {
			log.Printf("Error scanning tuteur row: %v", err)
			return nil, err
		}
		if tuteur.IdInterneMdt != "" {
			tuteurs = append(tuteurs, tuteur)
		}
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating tuteurs rows: %v", err)
		return nil, err
	}

	return tuteurs, nil
}

// GetEmployeur récupère les données d'employeur
func (r *OraclePersonnePhysiqueRepository) GetEmployeur(ctx context.Context, idInterneClt string) (*models.PersonnePhysiqueEmployeur, error) {
	query := `SELECT IdInterneEmpl,DenominationSociale,RCCM,NIF,NIFP,DateCreation,DateEntree 
	          FROM DIAMA.DBANK_EMPLOYEUR 
	          WHERE IdInterneClt = :1`

	row := r.db.QueryRowContext(ctx, query, idInterneClt)
	
	var employeur models.PersonnePhysiqueEmployeur
	if err := row.Scan(&employeur.IdInterneEmpl, &employeur.DenominationSociale, &employeur.RCCM, 
		&employeur.NIF, &employeur.NIFP, &employeur.DateCreation, &employeur.DateEntree); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Printf("Error scanning employeur for %s: %v", idInterneClt, err)
		return nil, err
	}

	return &employeur, nil
}

// GetAdditionnelles récupère les données additionnelles
func (r *OraclePersonnePhysiqueRepository) GetAdditionnelles(ctx context.Context, idInterneClt string) (*models.PersonnePhysiqueAdditionnelles, error) {
	query := `SELECT Cle,Valeur 
	          FROM DIAMA.DBANK_PP_ADDITIONNELLES 
	          WHERE IdInterneClt = :1`

	row := r.db.QueryRowContext(ctx, query, idInterneClt)
	
	var additionnelles models.PersonnePhysiqueAdditionnelles
	if err := row.Scan(&additionnelles.Cle, &additionnelles.Valeur); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Printf("Error scanning additionnelles for %s: %v", idInterneClt, err)
		return nil, err
	}

	return &additionnelles, nil
}

// ExecuteStoredProcedure exécute la procédure stockée pour générer les données
func (r *OraclePersonnePhysiqueRepository) ExecuteStoredProcedure(ctx context.Context, dateDeclaration string) error {
	query := `BEGIN DIAMA.GEN_PERSPHYSIQUE(:1); END;`

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error preparing stored procedure: %v", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, dateDeclaration)
	if err != nil {
		log.Printf("Error executing stored procedure: %v", err)
		return err
	}

	return nil
}
