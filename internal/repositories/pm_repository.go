package repositories

import (
	"context"
	"database/sql"
	"log"
	"sicbcrg.diamabank.com/internal/models"
)

// PersonneMoraleRepository définit les opérations de données pour PersonneMorale
type PersonneMoraleRepository interface {
	GetAll(ctx context.Context) ([]models.PersonneMorale, error)
	GetMandataires(ctx context.Context, idInterneClt string) ([]models.PersonneMoraleMandataire, error)
	GetComptesAssocies(ctx context.Context, idInterneClt string) ([]models.PersonneMoraleCompteAssocie, error)
	GetMandatairesAssocie(ctx context.Context, idInterneClt string) ([]models.PersonneMoraleMandataireMandataireAssocie, error)
	GetActionnaires(ctx context.Context, idInterneClt string) ([]models.PersonneMoraleActionnaire, error)
	GetAdditionnelles(ctx context.Context, idInterneClt string) ([]models.PersonneMoraleAdditionnelles, error)
	ExecuteStoredProcedure(ctx context.Context, dateDeclaration string) error
}

// OraclePersonneMoraleRepository implémente PersonneMoraleRepository pour Oracle
type OraclePersonneMoraleRepository struct {
	db *sql.DB
}

// NewOraclePersonneMoraleRepository crée une nouvelle instance
func NewOraclePersonneMoraleRepository(db *sql.DB) PersonneMoraleRepository {
	return &OraclePersonneMoraleRepository{db: db}
}

// GetAll récupère toutes les personnes morales
func (r *OraclePersonneMoraleRepository) GetAll(ctx context.Context) ([]models.PersonneMorale, error) {
	var personnes []models.PersonneMorale

	query := `SELECT * FROM DIAMA.DBANK_PERSMORALE`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		log.Printf("Error querying PersonneMorale: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var pm models.PersonneMorale
		if err := rows.Scan(&pm.Client, &pm.EstDeclare, &pm.DateDeclare, &pm.Datmaj, &pm.NatDec, &pm.NatClient, 
			&pm.IdInterneClt, &pm.DenomSocial, &pm.Sigle, &pm.DatCreat, &pm.Statut, &pm.DatCreaPart, &pm.FormeJuridique, 
			&pm.PaysSiegeSocial, &pm.VilleSiegeSocial, &pm.Mobile, &pm.Email, &pm.SiteWeb, &pm.Adress, &pm.CommuneAdress, 
			&pm.CodePostal, &pm.Resident, &pm.RCCM, &pm.NIF, &pm.NIFP, &pm.NumAgrement, &pm.NumSecSoc, &pm.SecActEcon, 
			&pm.SectInst, &pm.SitBancaire, &pm.DateDebIB, &pm.DateFinIB); err != nil {
			log.Printf("Error scanning PersonneMorale row: %v", err)
			return nil, err
		}
		personnes = append(personnes, pm)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating PersonneMorale rows: %v", err)
		return nil, err
	}

	return personnes, nil
}

// GetMandataires récupère les mandataires d'une personne morale
func (r *OraclePersonneMoraleRepository) GetMandataires(ctx context.Context, idInterneClt string) ([]models.PersonneMoraleMandataire, error) {
	var mandataires []models.PersonneMoraleMandataire

	query := `SELECT * FROM DIAMA.DBANK_MANDATAIRE WHERE IdInterneClt = :1`

	rows, err := r.db.QueryContext(ctx, query, idInterneClt)
	if err != nil {
		log.Printf("Error querying mandataires for %s: %v", idInterneClt, err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var mandataire models.PersonneMoraleMandataire
		if err := rows.Scan(&mandataire.Client, &mandataire.Idp, &mandataire.EstDeclare, &mandataire.DateDeclare, 
			&mandataire.Datmaj, &mandataire.IdInterneClt, &mandataire.IdInterneMdt, &mandataire.Qualite, 
			&mandataire.DatDebMdt, &mandataire.DatFinMdt); err != nil {
			log.Printf("Error scanning mandataire row: %v", err)
			return nil, err
		}
		mandataires = append(mandataires, mandataire)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating mandataires rows: %v", err)
		return nil, err
	}

	return mandataires, nil
}

// GetComptesAssocies récupère les comptes associés
func (r *OraclePersonneMoraleRepository) GetComptesAssocies(ctx context.Context, idInterneClt string) ([]models.PersonneMoraleCompteAssocie, error) {
	var comptes []models.PersonneMoraleCompteAssocie

	query := `SELECT * FROM DIAMA.DBANK_PERSMORALE_CPT WHERE IdInterneClt = :1`

	rows, err := r.db.QueryContext(ctx, query, idInterneClt)
	if err != nil {
		log.Printf("Error querying comptes associes for %s: %v", idInterneClt, err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var compte models.PersonneMoraleCompteAssocie
		if err := rows.Scan(&compte.Client, &compte.Compte, &compte.Devise, &compte.Ncg, &compte.EstDeclare, 
			&compte.DateDeclare, &compte.Datmaj, &compte.IdInterneClt, &compte.CodAgce, &compte.NumCpt, 
			&compte.CleRib, &compte.StatCpt); err != nil {
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

// GetMandatairesAssocie récupère les mandataires associés à un compte
func (r *OraclePersonneMoraleRepository) GetMandatairesAssocie(ctx context.Context, idInterneClt string) ([]models.PersonneMoraleMandataireMandataireAssocie, error) {
	var mandataires []models.PersonneMoraleMandataireMandataireAssocie

	query := `SELECT * FROM DIAMA.DBANK_MANDASSOCIE WHERE IdInterneClt = :1`

	rows, err := r.db.QueryContext(ctx, query, idInterneClt)
	if err != nil {
		log.Printf("Error querying mandataires associes for %s: %v", idInterneClt, err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var mandataire models.PersonneMoraleMandataireMandataireAssocie
		if err := rows.Scan(&mandataire.Client, &mandataire.Idp, &mandataire.EstDeclare, &mandataire.DateDeclare, 
			&mandataire.Datmaj, &mandataire.IdInterneClt, &mandataire.IdInterneMdtCpt, &mandataire.DatDebMdt, 
			&mandataire.DatFinMdt); err != nil {
			log.Printf("Error scanning mandataire associe row: %v", err)
			return nil, err
		}
		mandataires = append(mandataires, mandataire)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating mandataires associes rows: %v", err)
		return nil, err
	}

	return mandataires, nil
}

// GetActionnaires récupère les actionnaires
func (r *OraclePersonneMoraleRepository) GetActionnaires(ctx context.Context, idInterneClt string) ([]models.PersonneMoraleActionnaire, error) {
	var actionnaires []models.PersonneMoraleActionnaire

	query := `SELECT * FROM DIAMA.DBANK_ACTIONNAIRE WHERE IdInterneClt = :1`

	rows, err := r.db.QueryContext(ctx, query, idInterneClt)
	if err != nil {
		log.Printf("Error querying actionnaires for %s: %v", idInterneClt, err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var actionnaire models.PersonneMoraleActionnaire
		if err := rows.Scan(&actionnaire.Client, &actionnaire.Idp, &actionnaire.EstDeclare, &actionnaire.DateDeclare, 
			&actionnaire.Datmaj, &actionnaire.IdInterneClt, &actionnaire.IdInterneAct, &actionnaire.PartAct, 
			&actionnaire.DaEntrAct); err != nil {
			log.Printf("Error scanning actionnaire row: %v", err)
			return nil, err
		}
		actionnaires = append(actionnaires, actionnaire)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating actionnaires rows: %v", err)
		return nil, err
	}

	return actionnaires, nil
}

// GetAdditionnelles récupère les données additionnelles
func (r *OraclePersonneMoraleRepository) GetAdditionnelles(ctx context.Context, idInterneClt string) ([]models.PersonneMoraleAdditionnelles, error) {
	var additionnelles []models.PersonneMoraleAdditionnelles

	query := `SELECT * FROM DIAMA.DBANK_PM_ADDITIONNELLES WHERE IdInterneClt = :1`

	rows, err := r.db.QueryContext(ctx, query, idInterneClt)
	if err != nil {
		log.Printf("Error querying additionnelles for %s: %v", idInterneClt, err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var additionnelle models.PersonneMoraleAdditionnelles
		if err := rows.Scan(&additionnelle.Client, &additionnelle.EstDeclare, &additionnelle.DateDeclare, 
			&additionnelle.Datmaj, &additionnelle.IdInterneClt, &additionnelle.Cle, &additionnelle.Valeur); err != nil {
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
func (r *OraclePersonneMoraleRepository) ExecuteStoredProcedure(ctx context.Context, dateDeclaration string) error {
	query := `BEGIN DIAMA.GEN_PERSMORALE(:1); END;`

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
