package repositories

import (
	"context"
	"database/sql"
	"log"
	"sicbcrg.diamabank.com/internal/models"
)

// EncoursRepository définit les opérations de données pour Encours
type EncoursRepository interface {
	GetAll(ctx context.Context) ([]models.Encours, error)
	GetAdditionnelles(ctx context.Context, refIntEng string) ([]models.EncoursAdditionnelles, error)
	ExecuteStoredProcedure(ctx context.Context, dateDeclaration string) error
}

// OracleEncoursRepository implémente EncoursRepository pour Oracle
type OracleEncoursRepository struct {
	db *sql.DB
}

// NewOracleEncoursRepository crée une nouvelle instance
func NewOracleEncoursRepository(db *sql.DB) EncoursRepository {
	return &OracleEncoursRepository{db: db}
}

// GetAll récupère tous les encours
func (r *OracleEncoursRepository) GetAll(ctx context.Context) ([]models.Encours, error) {
	var encours []models.Encours

	query := `SELECT NatDec, RefIntEng, CodDev, DatEch, MntDerEch, MonPai, DatPai, MntHBil, MntRemAnt, MntCRDU, 
	          MntCreRat, MntUtilise, TO_CHAR(MntAgi), MntCapImp, MntTotImp, DatDefaill, MntPro, MntPerte, NbrEchPay, 
	          NbrEchImp, NbrEchRest, QualiCre, PD, LGD, CCF, IFRSStage, DatEvent 
	          FROM DIAMA.DBANK_ENCOURS`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		log.Printf("Error querying Encours: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var enc models.Encours
		if err := rows.Scan(&enc.NatDec, &enc.RefIntEng, &enc.CodDev, &enc.DatEch, &enc.MntDerEch, &enc.MonPai, 
			&enc.DatPai, &enc.MntHBil, &enc.MntRemAnt, &enc.MntCRDU, &enc.MntCreRat, &enc.MntUtilise, &enc.MntAgi, 
			&enc.MntCapImp, &enc.MntTotImp, &enc.DatDefaill, &enc.MntPro, &enc.MntPerte, &enc.NbrEchPay, &enc.NbrEchImp,
			&enc.NbrEchRest, &enc.QualiCre, &enc.PD, &enc.LGD, &enc.CCF, &enc.IFRSStage, &enc.DatEvent); err != nil {
			log.Printf("Error scanning Encours row: %v", err)
			return nil, err
		}
		encours = append(encours, enc)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating Encours rows: %v", err)
		return nil, err
	}

	return encours, nil
}

// GetAdditionnelles récupère les données additionnelles d'un encours
func (r *OracleEncoursRepository) GetAdditionnelles(ctx context.Context, refIntEng string) ([]models.EncoursAdditionnelles, error) {
	var additionnelles []models.EncoursAdditionnelles

	query := `SELECT Cle, Valeur FROM DIAMA.DBANK_ENC_ADDITIONNELLES WHERE RefIntEng = :1`

	rows, err := r.db.QueryContext(ctx, query, refIntEng)
	if err != nil {
		log.Printf("Error querying additionnelles for %s: %v", refIntEng, err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var additionnelle models.EncoursAdditionnelles
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
func (r *OracleEncoursRepository) ExecuteStoredProcedure(ctx context.Context, dateDeclaration string) error {
	query := `BEGIN DIAMA.GEN_ENCOURS(:1); END;`

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
