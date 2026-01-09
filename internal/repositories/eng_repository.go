package repositories

import (
	"context"
	"database/sql"
	"log"
	"sicbcrg.diamabank.com/internal/models"
)

// EngagementsRepository définit les opérations de données pour Engagements
type EngagementsRepository interface {
	GetAll(ctx context.Context) ([]models.Engagements, error)
	GetBeneficiaires(ctx context.Context, refIntEng string) ([]models.EngagementsBeneficiaire, error)
	GetAdditionnelles(ctx context.Context, refIntEng string) ([]models.EngagementsAdditionnelles, error)
	ExecuteStoredProcedure(ctx context.Context, dateDeclaration string) error
}

// OracleEngagementsRepository implémente EngagementsRepository pour Oracle
type OracleEngagementsRepository struct {
	db *sql.DB
}

// NewOracleEngagementsRepository crée une nouvelle instance
func NewOracleEngagementsRepository(db *sql.DB) EngagementsRepository {
	return &OracleEngagementsRepository{db: db}
}

// GetAll récupère tous les engagements
func (r *OracleEngagementsRepository) GetAll(ctx context.Context) ([]models.Engagements, error) {
	var engagements []models.Engagements

	query := `SELECT RefIntEng, NatDec, TypEve, LigneParent, RefIntLigne, RefDemandeEng, DatDem, TypModif, EstDout, 
	          Cloture, MotifCloture, DatClo, DatAccord, DateMEP, TypEng, TO_CHAR(MntEng), TO_CHAR(MntInt), CodDev, 
	          PeriodRemb, TxIntEng, TypTxInt, TxComm, IndRef, Sprd, TxEffGlob, MoyRemb, TypAmo, TypDiffAmo, UnitDur, PerDiffAmo,
	          MntEch, NbrEch, DatPremEch, DatFin, MntFrais, MntComm, CodAgce, EstRachatCreance,
	          ParCont, TO_CHAR(ValNom), TO_CHAR(ValCess), DatEvent, IdIntBen, PourBenef 
	          FROM DIAMA.DBANK_ENGAGEMENTS`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		log.Printf("Error querying Engagements: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var eng models.Engagements
		if err := rows.Scan(&eng.RefIntEng, &eng.NatDec, &eng.TypEve,
			&eng.LigneParent, &eng.RefIntLigne, &eng.RefDemandeEng, &eng.DatDem, &eng.TypeModif, &eng.EstDout, &eng.Cloture, 
			&eng.DatAccord, &eng.MotifCloture, &eng.DatClo, &eng.DateMEP, &eng.TypEng, &eng.MntEng, &eng.MntInt, &eng.CodDev, 
			&eng.PeriodRemb, &eng.TxIntEng, &eng.TypTxInt, &eng.TxComm, &eng.IndRef, &eng.Sprd, &eng.TxEffGlob, &eng.MoyRemb, 
			&eng.TypAmo, &eng.TypDiffAmo, &eng.UnitDur, &eng.PerDiffAmo, &eng.MntEch, &eng.NbrEch, &eng.DatPremEch, &eng.DatFin, 
			&eng.MntFrais, &eng.MntComm, &eng.CodAgce, &eng.EstRachatCreance, &eng.ParCont, &eng.ValNom, &eng.ValCess, 
			&eng.DatEvent, &eng.IdIntBen, &eng.PourBenef); err != nil {
			log.Printf("Error scanning Engagements row: %v", err)
			return nil, err
		}
		engagements = append(engagements, eng)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating Engagements rows: %v", err)
		return nil, err
	}

	return engagements, nil
}

// GetBeneficiaires récupère les bénéficiaires d'un engagement
func (r *OracleEngagementsRepository) GetBeneficiaires(ctx context.Context, refIntEng string) ([]models.EngagementsBeneficiaire, error) {
	var beneficiaires []models.EngagementsBeneficiaire

	// Les bénéficiaires sont inclus dans la table principale pour Engagements
	// Donc nous retournons seulement le bénéficiaire si disponible
	var beneficiaire models.EngagementsBeneficiaire
	query := `SELECT IdIntBen, PourBenef FROM DIAMA.DBANK_ENGAGEMENTS WHERE RefIntEng = :1`

	row := r.db.QueryRowContext(ctx, query, refIntEng)
	if err := row.Scan(&beneficiaire.IdIntBen, &beneficiaire.PourBenef); err != nil {
		if err == sql.ErrNoRows {
			return beneficiaires, nil
		}
		log.Printf("Error scanning beneficiaire for %s: %v", refIntEng, err)
		return nil, err
	}

	beneficiaires = append(beneficiaires, beneficiaire)
	return beneficiaires, nil
}

// GetAdditionnelles récupère les données additionnelles d'un engagement
func (r *OracleEngagementsRepository) GetAdditionnelles(ctx context.Context, refIntEng string) ([]models.EngagementsAdditionnelles, error) {
	var additionnelles []models.EngagementsAdditionnelles

	query := `SELECT Cle, Valeur FROM DIAMA.DBANK_ENG_ADDITIONNELLES WHERE RefIntEng = :1`

	rows, err := r.db.QueryContext(ctx, query, refIntEng)
	if err != nil {
		log.Printf("Error querying additionnelles for %s: %v", refIntEng, err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var additionnelle models.EngagementsAdditionnelles
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
func (r *OracleEngagementsRepository) ExecuteStoredProcedure(ctx context.Context, dateDeclaration string) error {
	query := `BEGIN DIAMA.GEN_ENGAGEMENTS(:1); END;`

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
