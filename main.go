package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/godror/godror"
	"github.com/google/uuid"
	"sicbcrg.diamabank.com/internal/models"
)

var NumDeclarationPP int = 0
var NumDeclarationPM int = 0
var NumDeclarationENG int = 0
var NumDeclarationENC int = 0
var NumDeclarationDEB int = 0

var db *sql.DB = nil
var err error
var headerXml string = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>` + "\n"

func main() {
	// Configuration depuis variables d'environnement
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "SICPROD"
	}
	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "NMu6F0DtoXFnnkyHt80Z"
	}
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "10.0.16.3"
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "1521"
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "ORCLPDB"
	}

	dataSourceName := fmt.Sprintf(`user="%s" password="%s" connectString="%s:%s/%s"`, dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err = sql.Open("godror", dataSourceName)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	fmt.Println("Successfully connected to Oracle Database!")

	//API
	router := http.NewServeMux()
	router.HandleFunc("/declaration/personnephysique", HandlerPersonnePhysique)
	router.HandleFunc("/declaration/personnemorale", HandlerPersonneMorale)
	router.HandleFunc("/declaration/engagements", HandlerEngagement)
	router.HandleFunc("/declaration/encours", HandlerEncours)
	router.HandleFunc("/declaration/comptedebiteurs", HandlerCompteDebiteurs)

	// Configuration serveur depuis variables d'environnement
	apiPort := os.Getenv("API_PORT")
	if apiPort == "" {
		apiPort = "8080"
	}
	serverAddr := "0.0.0.0:" + apiPort

	server := http.Server{
		Addr:         serverAddr,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  10 * time.Second,
		Handler:      router,
	}

	fmt.Printf("API Server listening on %s\n", serverAddr)
	log.Fatal(server.ListenAndServe())
}

// Personne Physique - OK
func HandlerPersonnePhysique(w http.ResponseWriter, r *http.Request) {
	var arr_pp []models.PersonnePhysique
	var pp models.PersonnePhysique

	var dateDeclaration string = r.URL.Query().Get("date")
	if dateDeclaration == "" {
		w.Write([]byte("Saisir une date valide"))
		return
	}

	//Exécution Procédure stockée
	query := `BEGIN DIAMA.GEN_PERSPHYSIQUE(:1); END;`
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	fmt.Printf("Date de la déclaration : %s\n", dateDeclaration)
	_, err = stmt.ExecContext(ctx, dateDeclaration)
	if err != nil {
		log.Fatal(err)
	}

	NumDeclarationPP++

	//Collecte des données pour formuler la réponse
	rows, err := db.Query("SELECT Client,Idp,EstDeclare,DateDeclare,Datmaj,NatDec,NatClient,NIN,IdInterneClt,DatCreaPart,NomNaiClt,NomMtlClt,PrenomClt,Sexe,DatNai,EtatCivil,NomPere,PrenomPere,NomNaiMere,PrmMre,VilleNai,PaysNai,NatClt,Resident,PaysRes,Mobile,Email,Adress,CommuneAdress,CodePostal,Profession,SecActEcon,SectInst,STutelle,StatutClt,DateDeces,SitBancaire,DateDebIB,DateFinIB FROM DIAMA.DBANK_PERSPHYSIQUE")
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&pp.Client, &pp.Idp, &pp.EstDeclare, &pp.DateDeclare, &pp.Datmaj,
			&pp.NatDec, &pp.NatClient, &pp.NIN, &pp.IdInterneClt, &pp.DatCreaPart, &pp.NomNaiClt, &pp.NomMtlClt,
			&pp.PrenomClt, &pp.Sexe, &pp.DatNai, &pp.EtatCivil, &pp.NomPere, &pp.PrenomPere, &pp.NomNaiMere,
			&pp.PrmMre, &pp.VilleNai, &pp.PaysNai, &pp.NatClt, &pp.Resident, &pp.PaysRes, &pp.Mobile,
			&pp.Email, &pp.Adress, &pp.CommuneAdress, &pp.CodePostal, &pp.Profession, &pp.SecActEcon, &pp.SectInst,
			&pp.STutelle, &pp.StatutClt, &pp.DateDeces, &pp.SitBancaire, &pp.DateDebIB, &pp.DateFinIB); err != nil {
			log.Fatalf("Error scanning row: %v\n\n", err)
		}

		//cas des comptes associes
		rowc, err := db.Query("SELECT CodAgce,NumCpt,CleRib,TypCpt,StatCpt FROM DIAMA.DBANK_PERSPHYSIQUE_CPT WHERE IdInterneClt = :1", pp.IdInterneClt)
		if err != nil {
			log.Fatalf("Error executing query for comptes associes: %v", err)
		}
		defer rowc.Close()
		var comptes []models.PersonnePhysiqueCompteAssocie
		for rowc.Next() {
			var compte models.PersonnePhysiqueCompteAssocie
			if err := rowc.Scan(&compte.CodAgce, &compte.NumCpt, &compte.CleRib, &compte.TypCpt, &compte.StatCpt); err != nil {
				log.Fatalf("Error scanning compte associe row: %v\n\n", err)
			}
			comptes = append(comptes, compte)
		}
		if len(comptes) > 0 {
			pp.CompteAssocie = &comptes
		}
		//end comptes associes

		//cas des pieces
		rowp, err := db.Query("SELECT TypPiece,NumPiece,DatEmiPiece,LieuEmiPiece,PaysEmiPiece,FinValPiece FROM DIAMA.DBANK_PP_PIECE WHERE IdInterneClt = :1", pp.IdInterneClt)
		if err != nil {
			log.Fatalf("Error executing query for pieces: %v", err)
		}
		defer rowp.Close()
		var pieces []models.PersonnePhysiquePiece
		for rowp.Next() {
			var piece models.PersonnePhysiquePiece
			if err := rowp.Scan(&piece.TypPiece, &piece.NumPiece, &piece.DatEmiPiece, &piece.LieuEmiPiece, &piece.PaysEmiPiece, &piece.FinValPiece); err != nil {
				log.Fatalf("Error scanning piece row: %v\n\n", err)
			}
			pieces = append(pieces, piece)
		}
		if len(pieces) > 0 {
			pp.Piece = &pieces
		}
		//end pieces

		//cas des donnees complementaires
		rowd, err := db.Query("SELECT NbPersCharge,RevMensMoy,DepMensMoy,PropLoc FROM DIAMA.DBANK_PP_COMPLEMENTAIRE WHERE IdInterneClt = :1", pp.IdInterneClt)
		if err != nil {
			log.Fatalf("Error executing query for donnees complementaires: %v", err)
		}
		defer rowd.Close()
		var donnee models.PersonnePhysiqueComplementaire
		if rowd.Next() {
			if err := rowd.Scan(&donnee.NbPersCharge, &donnee.RevMensMoy, &donnee.DepMensMoy, &donnee.PropLoc); err != nil {
				log.Fatalf("Error scanning donnee complementaire row: %v\n\n", err)
			}
			pp.DonneeComplementaire = &donnee
		}
		//end donnees complementaires

		//cas des tuteurs curateurs
		if pp.STutelle != "O" {
			rowt, err := db.Query("SELECT IdInterneMdt,Qualite,DatDbtMdt FROM DIAMA.DBANK_TUTEURCURATEUR WHERE IdInterneClt = :1", pp.IdInterneClt)
			if err != nil {
				log.Fatalf("Error executing query for tuteurs curateurs: %v", err)
			}
			defer rowt.Close()
			var tuteurs []models.PersonnePhysiqueTuteurCurateur
			for rowt.Next() {
				var tuteur models.PersonnePhysiqueTuteurCurateur
				if err := rowt.Scan(&tuteur.IdInterneMdt, &tuteur.Qualite, &tuteur.DatDbtMdt); err != nil {
					log.Fatalf("Error scanning tuteur curateur row: %v\n\n", err)
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
		rowe, err := db.Query("SELECT IdInterneEmpl,DenominationSociale,RCCM,NIF,NIFP,DateCreation,DateEntree  FROM DIAMA.DBANK_EMPLOYEUR WHERE IdInterneClt = :1", pp.IdInterneClt)
		if err != nil {
			log.Fatalf("Error executing query for employeur: %v", err)
		}
		defer rowe.Close()
		var employeur models.PersonnePhysiqueEmployeur
		if rowe.Next() {
			if err := rowe.Scan(&employeur.IdInterneEmpl, &employeur.DenominationSociale, &employeur.RCCM, &employeur.NIF, &employeur.NIFP, &employeur.DateCreation, &employeur.DateEntree); err != nil {
				log.Fatalf("Error scanning employeur row: %v\n\n", err)
			}
			pp.Employeur = &employeur
		}
		//end employeurs

		//cas des donnees additionnelles
		rowa, err := db.Query("SELECT Cle,Valeur FROM DIAMA.DBANK_PP_ADDITIONNELLES WHERE IdInterneClt = :1", pp.IdInterneClt)
		if err != nil {
			log.Fatalf("Error executing query for donnees additionnelles: %v", err)
		}
		defer rowa.Close()
		var additionnelles models.PersonnePhysiqueAdditionnelles
		if rowa.Next() {
			if err := rowa.Scan(&additionnelles.Cle, &additionnelles.Valeur); err != nil {
				log.Fatalf("Error scanning donnees additionnelles row: %v\n\n", err)
			}
			pp.DonneesAdditionnelles = &additionnelles
		}
		//end donnees additionnelles

		arr_pp = append(arr_pp, pp)
	}

	//Préparation de la structure de réponse
	declaration := DeclarationPP{
		// NumDec:           strconv.Itoa(NumDeclarationPP),
		NumDec:           fmt.Sprintf("%04d", NumDeclarationPP),
		PartEmtr:         "038",
		TypDec:           "01",
		NbrDec:           strconv.Itoa(len(arr_pp)),
		DateDec:          strings.Join(strings.Split(time.Now().Format("02-01-2006"), "-"), ""),
		PersonnePhysique: arr_pp,
	}
	response := Response{
		Meta: Meta{
			Status:      200,
			Message:     "success",
			RequestTime: time.Now().Format("02-01-2006 15:04:05"),
			RequestID:   uuid.NewString(),
		},
		Data: Data{
			Declaration: declaration,
		},
	}

	w.Header().Set("Content-Type", "application/xml;charset='utf-8'")
	xmlData, err := xml.Marshal(response)
	if err != nil {
		http.Error(w, "Erreur lors de l'encodage XML", http.StatusInternalServerError)
		return
	}
	// w.Write([]byte(headerXml))
	w.Write([]byte(headerXml))
	w.Write(xmlData)

	arr_pp = nil
}

// Personne Morale - OK
func HandlerPersonneMorale(w http.ResponseWriter, r *http.Request) {
	var arr_pm []models.PersonneMorale
	var pm models.PersonneMorale

	var dateDeclaration string = r.URL.Query().Get("date")
	if dateDeclaration == "" {
		w.Write([]byte("Saisir une date valide"))
		return
	}

	//Exécution Procédure stockée
	query := `BEGIN DIAMA.GEN_PERSMORALE(:1); END;`
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	fmt.Printf("Date de la déclaration : %s\n", dateDeclaration)
	_, err = stmt.ExecContext(ctx, dateDeclaration)
	if err != nil {
		log.Fatal(err)
	}

	NumDeclarationPM++

	//Collecte des données pour formuler la réponse
	rows, err := db.Query("SELECT * FROM DIAMA.DBANK_PERSMORALE")
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&pm.Client, &pm.EstDeclare, &pm.DateDeclare, &pm.Datmaj, &pm.NatDec, &pm.NatClient, &pm.IdInterneClt, &pm.DenomSocial,
			&pm.Sigle, &pm.DatCreat, &pm.Statut, &pm.DatCreaPart, &pm.FormeJuridique, &pm.PaysSiegeSocial, &pm.VilleSiegeSocial, &pm.Mobile, &pm.Email,
			&pm.SiteWeb, &pm.Adress, &pm.CommuneAdress, &pm.CodePostal, &pm.Resident, &pm.RCCM, &pm.NIF, &pm.NIFP, &pm.NumAgrement, &pm.NumSecSoc,
			&pm.SecActEcon, &pm.SectInst, &pm.SitBancaire, &pm.DateDebIB, &pm.DateFinIB); err != nil {
			log.Fatalf("Error 1 - scanning row: %v\n\n", err)
		}

		//cas Mandataires
		rowc, err := db.Query("SELECT * FROM DIAMA.DBANK_MANDATAIRE WHERE IdInterneClt = :1", pm.IdInterneClt)
		if err != nil {
			log.Fatalf("Error executing query for mandataire: %v", err)
		}
		defer rowc.Close()
		var mandataires []models.PersonneMoraleMandataire
		for rowc.Next() {
			var mandataire models.PersonneMoraleMandataire
			if err := rowc.Scan(&mandataire.Client, &mandataire.Idp, &mandataire.EstDeclare, &mandataire.DateDeclare, &mandataire.Datmaj, &mandataire.IdInterneClt,
				&mandataire.IdInterneMdt, &mandataire.Qualite, &mandataire.DatDebMdt, &mandataire.DatFinMdt); err != nil {
				log.Fatalf("Error 2 scanning mandaaire associe: %v\n\n", err)
			}
			mandataires = append(mandataires, mandataire)
		}
		if len(mandataires) > 0 {
			pm.Mandataires = &mandataires
		}
		//end mandataires

		//cas des comptes associes
		rowp, err := db.Query("SELECT * FROM DIAMA.DBANK_PERSMORALE_CPT WHERE IdInterneClt = :1", pm.IdInterneClt)
		if err != nil {
			log.Fatalf("Error executing query for compte associés: %v", err)
		}
		defer rowp.Close()
		var comptesassocies []models.PersonneMoraleCompteAssocie
		for rowp.Next() {
			var compteassocie models.PersonneMoraleCompteAssocie
			if err := rowp.Scan(&compteassocie.Client, &compteassocie.Compte, &compteassocie.Devise, &compteassocie.Ncg, &compteassocie.EstDeclare,
				&compteassocie.DateDeclare, &compteassocie.Datmaj, &compteassocie.IdInterneClt, &compteassocie.CodAgce, &compteassocie.NumCpt,
				&compteassocie.CleRib, &compteassocie.StatCpt); err != nil {
				log.Fatalf("Error 3 scanning piece row: %v\n\n", err)
			}

			compteassocie.MandataireAssocie = nil
			//cas mandataires associes au compte
			rowm, err := db.Query("SELECT * FROM DIAMA.DBANK_MANDASSOCIE WHERE IdInterneClt = :1", compteassocie.IdInterneClt)
			if err != nil {
				log.Fatalf("Error executing query for mandataires associes au compte: %v", err)
			}
			defer rowm.Close()
			var mandatairesassocies []models.PersonneMoraleMandataireMandataireAssocie
			for rowm.Next() {
				var mandataireassocie models.PersonneMoraleMandataireMandataireAssocie
				if err := rowm.Scan(&mandataireassocie.Client, &mandataireassocie.Idp, &mandataireassocie.EstDeclare, &mandataireassocie.DateDeclare,
					&mandataireassocie.Datmaj, &mandataireassocie.IdInterneClt, &mandataireassocie.IdInterneMdtCpt, &mandataireassocie.DatDebMdt,
					&mandataireassocie.DatFinMdt); err != nil {
					log.Fatalf("Error scanning mandataire associe au compte: %v\n\n", err)
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
		rowa, err := db.Query("SELECT * FROM DIAMA.DBANK_ACTIONNAIRE WHERE IdInterneClt = :1", pm.IdInterneClt)
		if err != nil {
			log.Fatalf("Error executing query for actionnaires: %v", err)
		}
		defer rowa.Close()
		var actionnaires []models.PersonneMoraleActionnaire
		for rowa.Next() {
			var actionnaire models.PersonneMoraleActionnaire
			if err := rowa.Scan(&actionnaire.Client, &actionnaire.Idp, &actionnaire.EstDeclare, &actionnaire.DateDeclare, &actionnaire.Datmaj,
				&actionnaire.IdInterneClt, &actionnaire.IdInterneAct, &actionnaire.PartAct, &actionnaire.DaEntrAct); err != nil {
				log.Fatalf("Error 4 scanning actionnaire row: %v\n\n", err)
			}
			actionnaires = append(actionnaires, actionnaire)
		}

		//end actionnaires
		if len(actionnaires) > 0 {
			pm.Actionnaire = &actionnaires
		}

		//cas des donnees additionnelles
		rowd, err := db.Query("SELECT * FROM DIAMA.DBANK_PM_ADDITIONNELLES WHERE IdInterneClt = :1", pm.IdInterneClt)
		if err != nil {
			log.Fatalf("Error executing query for donnees additionnelles: %v", err)
		}
		defer rowd.Close()
		var additionnelles []models.PersonneMoraleAdditionnelles
		for rowd.Next() {
			var additionnelle models.PersonneMoraleAdditionnelles
			if err := rowd.Scan(&additionnelle.Client, &additionnelle.EstDeclare, &additionnelle.DateDeclare, &additionnelle.Datmaj,
				&additionnelle.IdInterneClt, &additionnelle.Cle, &additionnelle.Valeur); err != nil {
				log.Fatalf("Error 5 scanning donnees additionnelles row: %v\n\n", err)
			}
			additionnelles = append(additionnelles, additionnelle)
		}
		if len(additionnelles) > 0 {
			pm.DonneesAdditionnelles = &additionnelles
		}
		//end donnees additionnelles

		arr_pm = append(arr_pm, pm)
	}

	//Préparation de la structure de réponse
	declaration := DeclarationPM{
		NumDec:         fmt.Sprintf("%04d", NumDeclarationPM),
		PartEmtr:       "038",
		TypDec:         "02",
		NbrDec:         strconv.Itoa(len(arr_pm)),
		DateDec:        strings.Join(strings.Split(time.Now().Format("02-01-2006"), "-"), ""),
		PersonneMorale: arr_pm,
	}
	response := Response{
		Meta: Meta{
			Status:      200,
			Message:     "success",
			RequestTime: time.Now().Format("02-01-2006 15:04:05"),
			RequestID:   uuid.NewString(),
		},
		Data: Data{
			Declaration: declaration,
		},
	}

	w.Header().Set("Content-Type", "application/xml;charset='utf-8'")
	fmt.Printf("Response: %+v\n", response)
	xmlData, err := xml.Marshal(response)
	if err != nil {
		http.Error(w, "Erreur lors de l'encodage XML", http.StatusInternalServerError)
		return
	}
	w.Write([]byte(headerXml))
	w.Write(xmlData)

	arr_pm = nil
}

// Engagements - OK
func HandlerEngagement(w http.ResponseWriter, r *http.Request) {
	arr_eng := []models.Engagements{}

	var dateDeclaration string = r.URL.Query().Get("date")
	if dateDeclaration == "" {
		w.Write([]byte("Saisir une date valide"))
		return
	}

	//Exécution Procédure stockée
	query := `BEGIN DIAMA.GEN_ENGAGEMENTS(:1); END;`
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	fmt.Printf("Date de la déclaration : %s\n", dateDeclaration)
	_, err = stmt.ExecContext(ctx, dateDeclaration)
	if err != nil {
		log.Fatal(err)
	}

	NumDeclarationENG++

	rows, err := db.Query(`SELECT RefIntEng, NatDec, TypEve, LigneParent, RefIntLigne, RefDemandeEng, DatDem, TypModif, EstDout, 
			Cloture, MotifCloture, DatClo,DatAccord, DateMEP, TypEng, TO_CHAR(MntEng), TO_CHAR(MntInt), CodDev, 
			PeriodRemb, TxIntEng, TypTxInt, TxComm, IndRef, Sprd, TxEffGlob, MoyRemb, TypAmo, TypDiffAmo, UnitDur, PerDiffAmo,
			MntEch, NbrEch, DatPremEch, DatFin, MntFrais, MntComm, CodAgce, EstRachatCreance,
			ParCont, TO_CHAR(ValNom), TO_CHAR(ValCess), DatEvent, IdIntBen, PourBenef FROM DIAMA.DBANK_ENGAGEMENTS`)
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var eng models.Engagements
		if err := rows.Scan(&eng.RefIntEng, &eng.NatDec, &eng.TypEve,
			&eng.LigneParent, &eng.RefIntLigne, &eng.RefDemandeEng, &eng.DatDem, &eng.TypeModif, &eng.EstDout, &eng.Cloture, &eng.DatAccord, &eng.MotifCloture, &eng.DatClo,
			&eng.DateMEP, &eng.TypEng, &eng.MntEng, &eng.MntInt, &eng.CodDev, &eng.PeriodRemb, &eng.TxIntEng, &eng.TypTxInt,
			&eng.TxComm, &eng.IndRef, &eng.Sprd, &eng.TxEffGlob, &eng.MoyRemb, &eng.TypAmo, &eng.TypDiffAmo, &eng.UnitDur, &eng.PerDiffAmo,
			&eng.MntEch, &eng.NbrEch, &eng.DatPremEch, &eng.DatFin, &eng.MntFrais, &eng.MntComm, &eng.CodAgce, &eng.EstRachatCreance,
			&eng.ParCont, &eng.ValNom, &eng.ValCess, &eng.DatEvent, &eng.IdIntBen, &eng.PourBenef); err != nil {
			log.Fatalf("Error scanning row: %v\n\n", err)
		}

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

		rowa, err := db.Query("SELECT Cle, Valeur FROM DIAMA.DBANK_ENG_ADDITIONNELLES WHERE RefIntEng = :1", eng.RefIntEng)
		if err != nil {
			log.Fatalf("Error executing query for donnees additionnelles: %v", err)
		}
		defer rowa.Close()

		for rowa.Next() {
			if err := rowa.Scan(&additionnelle.Cle, &additionnelle.Valeur); err != nil {
				log.Fatalf("Error scanning donnees additionnelles row: %v\n\n", err)
			}
			additionnelles = append(additionnelles, additionnelle)
		}
		if len(additionnelles) > 0 {
			eng.DonneesAdditionnelles = &additionnelles
		}
		// end données additionnelles

		//Ajout de l'engagement à la liste
		arr_eng = append(arr_eng, eng)
	}

	//Préparation de la structure de réponse
	declaration := DeclarationENG{
		NumDec:     fmt.Sprintf("%04d", NumDeclarationENG),
		PartEmtr:   "038",
		TypDec:     "01",
		NbrDec:     strconv.Itoa(len(arr_eng)),
		DateDec:    strings.Join(strings.Split(time.Now().Format("02-01-2006"), "-"), ""),
		Engagement: arr_eng,
	}
	response := Response{
		Meta: Meta{
			Status:      200,
			Message:     "success",
			RequestTime: time.Now().Format("02-01-2006 15:04:05"),
			RequestID:   uuid.NewString(),
		},
		Data: Data{
			Declaration: declaration,
		},
	}

	w.Header().Set("Content-Type", "application/xml;charset='utf-8'")
	xmlData, err := xml.Marshal(response)
	if err != nil {
		http.Error(w, "Erreur lors de l'encodage XML", http.StatusInternalServerError)
		return
	}
	w.Write([]byte(headerXml))
	w.Write(xmlData)

	arr_eng = nil
}

// Encours - OK
func HandlerEncours(w http.ResponseWriter, r *http.Request) {
	var arr_enc []models.Encours

	var dateDeclaration string = r.URL.Query().Get("date")
	if dateDeclaration == "" {
		w.Write([]byte("Saisir une date valide"))
		return
	}

	//Exécution Procédure stockée
	query := `BEGIN DIAMA.GEN_ENCOURS(:1); END;`
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	fmt.Printf("Date de la déclaration : %s\n", dateDeclaration)
	_, err = stmt.ExecContext(ctx, dateDeclaration)
	if err != nil {
		log.Fatal(err)
	}

	//À implémenter
	//Collecte des données pour formuler la réponse
	rows, err := db.Query(`SELECT NatDec,RefIntEng,CodDev,DatEch,MntDerEch,MonPai,DatPai,MntHBil,MntRemAnt,MntCRDU,MntCreRat,MntUtilise,TO_CHAR(MntAgi),MntCapImp,
    MntTotImp, DatDefaill, MntPro, MntPerte, NbrEchPay, NbrEchImp, NbrEchRest, QualiCre, PD, LGD, CCF, IFRSStage, DatEvent FROM DIAMA.DBANK_ENCOURS`)
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}
	defer rows.Close()

	NumDeclarationENC++
	for rows.Next() {
		var enc models.Encours
		if err := rows.Scan(&enc.NatDec, &enc.RefIntEng, &enc.CodDev, &enc.DatEch, &enc.MntDerEch, &enc.MonPai, &enc.DatPai, &enc.MntHBil, &enc.MntRemAnt, &enc.MntCRDU,
			&enc.MntCreRat, &enc.MntUtilise, &enc.MntAgi, &enc.MntCapImp, &enc.MntTotImp, &enc.DatDefaill, &enc.MntPro, &enc.MntPerte, &enc.NbrEchPay, &enc.NbrEchImp,
			&enc.NbrEchRest, &enc.QualiCre, &enc.PD, &enc.LGD, &enc.CCF, &enc.IFRSStage, &enc.DatEvent); err != nil {
			log.Fatalf("Error scanning row: %v\n\n", err)
		}

		// Données additionnelles
		var additionnelles []models.EncoursAdditionnelles
		var additionnelle models.EncoursAdditionnelles

		rowa, err := db.Query("SELECT Cle, Valeur FROM DIAMA.DBANK_ENC_ADDITIONNELLES WHERE RefIntEng = :1", enc.RefIntEng)
		if err != nil {
			log.Fatalf("Error executing query for donnees additionnelles: %v", err)
		}
		defer rowa.Close()

		for rowa.Next() {
			if err := rowa.Scan(&additionnelle.Cle, &additionnelle.Valeur); err != nil {
				log.Fatalf("Error scanning donnees additionnelles row: %v\n\n", err)
			}
			additionnelles = append(additionnelles, additionnelle)
		}
		if len(additionnelles) > 0 {
			enc.DonneesAdditionnelles = &additionnelles
		}
		// end données additionnelles

		//Ajout de l'encours à la liste
		arr_enc = append(arr_enc, enc)
	}
	date, err := time.Parse("02/01/06", dateDeclaration)
	if err != nil {
		log.Fatalf("Error parsing date: %v", err)
	}
	//réponse
	declaration := DeclarationENC{
		NumDec:            fmt.Sprintf("%04d", NumDeclarationENC),
		PartEmtr:          "038",
		TypDec:            "52",
		NbrDec:            strconv.Itoa(len(arr_enc)),
		DateDec:           strings.Join(strings.Split(time.Now().Format("02-01-2006"), "-"), ""),
		DatArr:            date.Format("020106"),
		EncoursEngagement: arr_enc,
	}
	response := Response{
		Meta: Meta{
			Status:      200,
			Message:     "success",
			RequestTime: time.Now().Format("02-01-2006 15:04:05"),
			RequestID:   uuid.NewString(),
		},
		Data: Data{
			Declaration: declaration,
		},
	}

	w.Header().Set("Content-Type", "application/xml;charset='utf-8'")
	xmlData, err := xml.Marshal(response)
	if err != nil {
		http.Error(w, "Erreur lors de l'encodage XML", http.StatusInternalServerError)
		return
	}
	w.Write([]byte(headerXml))
	w.Write(xmlData)

	arr_enc = nil
}

// Compte Débiteurs - OK
func HandlerCompteDebiteurs(w http.ResponseWriter, r *http.Request) {
	var arr_deb []models.CompteDebiteurs
	var dateDeclaration string = r.URL.Query().Get("date")
	if dateDeclaration == "" {
		w.Write([]byte("Saisir une date valide"))
		return
	}

	//Exécution Procédure stockée
	query := `BEGIN DIAMA.GEN_CPTDEBITEURS(:1); END;`
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	fmt.Printf("Date de la déclaration : %s\n", dateDeclaration)
	_, err = stmt.ExecContext(ctx, dateDeclaration)
	if err != nil {
		log.Fatal(err)
	}

	NumDeclarationDEB++

	query = `SELECT NatDec,CodDev,Rib,TO_CHAR(SoldeDeb),DateDefaill,NbrJourDebMax,TO_CHAR(SoldeDebMax),TO_CHAR(MntProv),TO_CHAR(MntPerte),
			TO_CHAR(MntAgi),TO_CHAR(QualiCre),IdIntTit	FROM DIAMA.DBANK_CPTDEBITEURS`

	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var deb models.CompteDebiteurs
		if err := rows.Scan(&deb.NatDec, &deb.CodDev, &deb.Rib, &deb.SoldeDeb, &deb.DateDefaill, &deb.NbrJourDebMax, &deb.SoldeDebMax,
			&deb.MntProv, &deb.MntPerte, &deb.MntAgi, &deb.QualiCre, &deb.IdIntTit); err != nil {
			log.Fatalf("Error scanning row: %v\n\n", err)
		}

		var titulaire models.CompteDebiteursTitulaire
		titulaire.IdIntTit = deb.IdIntTit
		deb.Titulaire = titulaire

		var additionnelles []models.CompteDebiteursAdditionnelles
		var additionnelle models.CompteDebiteursAdditionnelles

		// Données additionnelles du titulaire
		rowa, err := db.Query("SELECT Cle, Valeur FROM DIAMA.DBANK_DEB_ADDITIONNELLES WHERE RefIntEng = :1", deb.IdIntTit)
		if err != nil {
			log.Fatalf("Error executing query for donnees additionnelles: %v", err)
		}
		defer rowa.Close()

		for rowa.Next() {
			if err := rowa.Scan(&additionnelle.Cle, &additionnelle.Valeur); err != nil {
				log.Fatalf("Error scanning donnees additionnelles row: %v\n\n", err)
			}
			additionnelles = append(additionnelles, additionnelle)
		}
		if len(additionnelles) > 0 {
			deb.DonneesAdditionnelles = &additionnelles
		}
		// end données additionnelles

		arr_deb = append(arr_deb, deb)
	}

	date, err := time.Parse("02/01/06", dateDeclaration)
	if err != nil {
		log.Fatalf("Error parsing date: %v", err)
	}
	//Préparation de la structure de réponse
	declaration := DeclarationDEB{
		NumDec:          fmt.Sprintf("%04d", NumDeclarationDEB),
		PartEmtr:        "038",
		TypDec:          "12",
		NbrDec:          strconv.Itoa(len(arr_deb)),
		DateDec:         strings.Join(strings.Split(time.Now().Format("02-01-2006"), "-"), ""),
		DatArr:          date.Format("020106"),
		CompteDebiteurs: arr_deb,
	}
	response := Response{
		Meta: Meta{
			Status:      200,
			Message:     "success",
			RequestTime: time.Now().Format("02-01-2006 15:04:05"),
			RequestID:   uuid.NewString(),
		},
		Data: Data{
			Declaration: declaration,
		},
	}

	w.Header().Set("Content-Type", "application/xml;charset='utf-8'")
	xmlData, err := xml.Marshal(response)
	if err != nil {
		http.Error(w, "Erreur lors de l'encodage XML", http.StatusInternalServerError)
		return
	}
	w.Write([]byte(headerXml))
	w.Write(xmlData)

	arr_deb = nil
}

type DeclarationPP struct {
	NumDec           string                    `xml:"NumDec,attr"`
	PartEmtr         string                    `xml:"PartEmtr,attr"`
	TypDec           string                    `xml:"TypDec,attr"`
	NbrDec           string                    `xml:"NbrDec,attr"`
	DateDec          string                    `xml:"DateDec,attr"`
	PersonnePhysique []models.PersonnePhysique `xml:"PersonnePhysique"`
}
type DeclarationPM struct {
	NumDec         string                  `xml:"NumDec,attr"`
	PartEmtr       string                  `xml:"PartEmtr,attr"`
	TypDec         string                  `xml:"TypDec,attr"`
	NbrDec         string                  `xml:"NbrDec,attr"`
	DateDec        string                  `xml:"DateDec,attr"`
	PersonneMorale []models.PersonneMorale `xml:"PersonneMorale"`
}

type DeclarationENG struct {
	NumDec     string               `xml:"NumDec,attr"`
	PartEmtr   string               `xml:"PartEmtr,attr"`
	TypDec     string               `xml:"TypDec,attr"`
	NbrDec     string               `xml:"NbrDec,attr"`
	DateDec    string               `xml:"DateDec,attr"`
	Engagement []models.Engagements `xml:"Engagement"`
}

type DeclarationENC struct {
	NumDec            string           `xml:"NumDec,attr"`
	PartEmtr          string           `xml:"PartEmtr,attr"`
	TypDec            string           `xml:"TypDec,attr"`
	NbrDec            string           `xml:"NbrDec,attr"`
	DateDec           string           `xml:"DateDec,attr"`
	DatArr            string           `xml:"DatArr,attr"`
	EncoursEngagement []models.Encours `xml:"EncoursEngagement"`
}

type DeclarationDEB struct {
	NumDec          string                   `xml:"NumDec,attr"`
	PartEmtr        string                   `xml:"PartEmtr,attr"`
	TypDec          string                   `xml:"TypDec,attr"`
	NbrDec          string                   `xml:"NbrDec,attr"`
	DateDec         string                   `xml:"DateDec,attr"`
	DatArr          string                   `xml:"DatArr,attr"`
	CompteDebiteurs []models.CompteDebiteurs `xml:"CompteDebiteur"`
}

// Define Data and Response structs at package level
type DataPP struct {
	Declaration DeclarationPP `xml:"declaration"`
}

type DataPM struct {
	Declaration DeclarationPM `xml:"declaration"`
}

type DataENG struct {
	Declaration DeclarationENG `xml:"declaration"`
}

type DataENC struct {
	Declaration DeclarationENG `xml:"declaration"`
}
type DataDEB struct {
	Declaration DeclarationDEB `xml:"declaration"`
}

type Data struct {
	Declaration interface{} `xml:"declaration"`
}

type Meta struct {
	Status      int    `xml:"status"`
	Message     string `xml:"message"`
	RequestTime string `xml:"requestTime"`
	RequestID   string `xml:"requestId"`
}

type Response struct {
	Meta Meta `xml:"meta"`
	Data Data `xml:"data"`
}
