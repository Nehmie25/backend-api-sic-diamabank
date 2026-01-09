package repositories

import (
	"context"
	"database/sql"
	"log"
	"sicbcrg.diamabank.com/internal/models"
)

// CompteDebiteurRepository définit les opérations de données pour CompteDebiteurs
type CompteDebiteurRepository interface {
	GetAll(ctx context.Context) ([]models.CompteDebiteurs, error)
	GetAdditionnelles(ctx context.Context, idIntTit string) ([]models.CompteDebiteursAdditionnelles, error)
	ExecuteStoredProcedure(ctx context.Context, dateDeclaration string) error
}

// OracleCompteDebiteurRepository implémente CompteDebiteurRepository pour Oracle
type OracleCompteDebiteurRepository struct {
	db *sql.DB
}

// NewOracleCompteDebiteurRepository crée une nouvelle instance
func NewOracleCompteDebiteurRepository(db *sql.DB) CompteDebiteurRepository {
	return &OracleCompteDebiteurRepository{db: db}
}

// GetAll récupère tous les comptes débiteurs
func (r *OracleCompteDebiteurRepository) GetAll(ctx context.Context) ([]models.CompteDebiteurs, error) {
	var comptes []models.CompteDebiteurs

	query := `SELECT NatDec, CodDev, Rib, TO_CHAR(SoldeDeb), DateDefaill, NbrJourDebMax, TO_CHAR(SoldeDebMax), 
	          TO_CHAR(MntProv), TO_CHAR(MntPerte), TO_CHAR(MntAgi), TO_CHAR(QualiCre), IdIntTit 
	          FROM DIAMA.DBANK_CPTDEBITEURS`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		log.Printf("Error querying CompteDebiteurs: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var deb models.CompteDebiteurs
		if err := rows.Scan(&deb.NatDec, &deb.CodDev, &deb.Rib, &deb.SoldeDeb, &deb.DateDefaill, &deb.NbrJourDebMax, 
			&deb.SoldeDebMax, &deb.MntProv, &deb.MntPerte, &deb.MntAgi, &deb.QualiCre, &deb.IdIntTit); err != nil {
			log.Printf("Error scanning CompteDebiteur row: %v", err)
			return nil, err
		}
		comptes = append(comptes, deb)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating CompteDebiteurs rows: %v", err)
		return nil, err
	}

	return comptes, nil
}

// GetAdditionnelles récupère les données additionnelles d'un compte débiteur
func (r *OracleCompteDebiteurRepository) GetAdditionnelles(ctx context.Context, idIntTit string) ([]models.CompteDebiteursAdditionnelles, error) {
	var additionnelles []models.CompteDebiteursAdditionnelles

	query := `SELECT Cle, Valeur FROM DIAMA.DBANK_DEB_ADDITIONNELLES WHERE RefIntEng = :1`

	rows, err := r.db.QueryContext(ctx, query, idIntTit)
	if err != nil {
		log.Printf("Error querying additionnelles for %s: %v", idIntTit, err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var additionnelle models.CompteDebiteursAdditionnelles
		if err := rows.Scan(&additionnelle.Cle, &additionnelle.Valeur); err != nil {
			log.Printf("Error scanning additionnelle row: %v", err)
			return nil, err
		}
		additionnelles = append(additionnelles, additionnelle)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating additionnelles rows: %v", err)
		return nil, err
	}

	return additionnelles, nil
}

// ExecuteStoredProcedure exécute la procédure stockée
func (r *OracleCompteDebiteurRepository) ExecuteStoredProcedure(ctx context.Context, dateDeclaration string) error {
	query := `BEGIN DIAMA.GEN_CPTDEBITEURS(:1); END;`

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
