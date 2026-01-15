package handler

import (
	"encoding/xml"
	"net/http"

	// "fmt"
	// "log"
	// "net/http"
	// "strconv"
	// "strings"
	// "time"

	// "github.com/google/uuid"
	"GoWebapitest/config"

	// "sicbcrg.diamabank.com/internal/adapter/repository"
	// "sicbcrg.diamabank.com/internal/adapter/repository/Oracle"
	// "sicbcrg.diamabank.com/internal/core/domain/models"
	"GoWebapitest/internal/core/service"
)

func HandlerPersonnePhysique(service *service.PersonnePhysiqueService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		date := r.URL.Query().Get("date")
		if date == "" {
			http.Error(w, "date requise", http.StatusBadRequest)
			return
		}

		declaration, err := service.BuildDeclarationepersonnePhysique(date)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := config.Response{
			Meta: config.Meta{
				Status:  200,
				Message: "success",
			},
			Data: config.Data{Declaration: declaration},
		}

		config.EnableCORS(w)
		w.Header().Set("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(response)
	}
}

// func HandlerPersonnePhysique(w http.ResponseWriter, r *http.Request) {
// 	var arr_pp []models.PersonnePhysique
// 	var pp models.PersonnePhysique
// 	var NumDeclarationPP int = 0
// 	var headerXml string = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>` + "\n"

// 	var dateDeclaration string = r.URL.Query().Get("date")
// 	if dateDeclaration == "" {
// 		w.Write([]byte("Saisir une date valide"))
// 		return
// 	}

// 	db:=oracle.Db()

// 	// //Exécution Procédure stockée
// 	// query := `BEGIN DIAMA.GEN_PERSPHYSIQUE(:1); END;`
// 	// ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
// 	// defer cancel()

// 	// stmt, err := db.PrepareContext(ctx, query)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// defer stmt.Close()

// 	// fmt.Printf("Date de la déclaration : %s\n", dateDeclaration)
// 	// _, err = stmt.ExecContext(ctx, dateDeclaration)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }

// 	repository.GEN_PERSPHYSIQUE(dateDeclaration,db)

// 	NumDeclarationPP++

// 	//Collecte des données pour formuler la réponse
// 	rows, err := db.Query("SELECT Client,Idp,EstDeclare,DateDeclare,Datmaj,NatDec,NatClient,NIN,IdInterneClt,DatCreaPart,NomNaiClt,NomMtlClt,PrenomClt,Sexe,DatNai,EtatCivil,NomPere,PrenomPere,NomNaiMere,PrmMre,VilleNai,PaysNai,NatClt,Resident,PaysRes,Mobile,Email,Adress,CommuneAdress,CodePostal,Profession,SecActEcon,SectInst,STutelle,StatutClt,DateDeces,SitBancaire,DateDebIB,DateFinIB FROM DIAMA.DBANK_PERSPHYSIQUE")
// 	if err != nil {
// 		log.Fatalf("Error executing query: %v", err)
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		if err := rows.Scan(&pp.Client, &pp.Idp, &pp.EstDeclare, &pp.DateDeclare, &pp.Datmaj,
// 			&pp.NatDec, &pp.NatClient, &pp.NIN, &pp.IdInterneClt, &pp.DatCreaPart, &pp.NomNaiClt, &pp.NomMtlClt,
// 			&pp.PrenomClt, &pp.Sexe, &pp.DatNai, &pp.EtatCivil, &pp.NomPere, &pp.PrenomPere, &pp.NomNaiMere,
// 			&pp.PrmMre, &pp.VilleNai, &pp.PaysNai, &pp.NatClt, &pp.Resident, &pp.PaysRes, &pp.Mobile,
// 			&pp.Email, &pp.Adress, &pp.CommuneAdress, &pp.CodePostal, &pp.Profession, &pp.SecActEcon, &pp.SectInst,
// 			&pp.STutelle, &pp.StatutClt, &pp.DateDeces, &pp.SitBancaire, &pp.DateDebIB, &pp.DateFinIB); err != nil {
// 			log.Fatalf("Error scanning row: %v\n\n", err)
// 		}

// 		//cas des comptes associes
// 		rowc, err := db.Query("SELECT CodAgce,NumCpt,CleRib,TypCpt,StatCpt FROM DIAMA.DBANK_PERSPHYSIQUE_CPT WHERE IdInterneClt = :1", pp.IdInterneClt)
// 		if err != nil {
// 			log.Fatalf("Error executing query for comptes associes: %v", err)
// 		}
// 		defer rowc.Close()
// 		var comptes []models.PersonnePhysiqueCompteAssocie
// 		for rowc.Next() {
// 			var compte models.PersonnePhysiqueCompteAssocie
// 			if err := rowc.Scan(&compte.CodAgce, &compte.NumCpt, &compte.CleRib, &compte.TypCpt, &compte.StatCpt); err != nil {
// 				log.Fatalf("Error scanning compte associe row: %v\n\n", err)
// 			}
// 			comptes = append(comptes, compte)
// 		}
// 		if len(comptes) > 0 {
// 			pp.CompteAssocie = &comptes
// 		}
// 		//end comptes associes

// 		//cas des pieces
// 		rowp, err := db.Query("SELECT TypPiece,NumPiece,DatEmiPiece,LieuEmiPiece,PaysEmiPiece,FinValPiece FROM DIAMA.DBANK_PP_PIECE WHERE IdInterneClt = :1", pp.IdInterneClt)
// 		if err != nil {
// 			log.Fatalf("Error executing query for pieces: %v", err)
// 		}
// 		defer rowp.Close()
// 		var pieces []models.PersonnePhysiquePiece
// 		for rowp.Next() {
// 			var piece models.PersonnePhysiquePiece
// 			if err := rowp.Scan(&piece.TypPiece, &piece.NumPiece, &piece.DatEmiPiece, &piece.LieuEmiPiece, &piece.PaysEmiPiece, &piece.FinValPiece); err != nil {
// 				log.Fatalf("Error scanning piece row: %v\n\n", err)
// 			}
// 			pieces = append(pieces, piece)
// 		}
// 		if len(pieces) > 0 {
// 			pp.Piece = &pieces
// 		}
// 		//end pieces

// 		//cas des donnees complementaires
// 		rowd, err := db.Query("SELECT NbPersCharge,RevMensMoy,DepMensMoy,PropLoc FROM DIAMA.DBANK_PP_COMPLEMENTAIRE WHERE IdInterneClt = :1", pp.IdInterneClt)
// 		if err != nil {
// 			log.Fatalf("Error executing query for donnees complementaires: %v", err)
// 		}
// 		defer rowd.Close()
// 		var donnee models.PersonnePhysiqueComplementaire
// 		if rowd.Next() {
// 			if err := rowd.Scan(&donnee.NbPersCharge, &donnee.RevMensMoy, &donnee.DepMensMoy, &donnee.PropLoc); err != nil {
// 				log.Fatalf("Error scanning donnee complementaire row: %v\n\n", err)
// 			}
// 			pp.DonneeComplementaire = &donnee
// 		}
// 		//end donnees complementaires

// 		//cas des tuteurs curateurs
// 		if pp.STutelle != "O" {
// 			rowt, err := db.Query("SELECT IdInterneMdt,Qualite,DatDbtMdt FROM DIAMA.DBANK_TUTEURCURATEUR WHERE IdInterneClt = :1", pp.IdInterneClt)
// 			if err != nil {
// 				log.Fatalf("Error executing query for tuteurs curateurs: %v", err)
// 			}
// 			defer rowt.Close()
// 			var tuteurs []models.PersonnePhysiqueTuteurCurateur
// 			for rowt.Next() {
// 				var tuteur models.PersonnePhysiqueTuteurCurateur
// 				if err := rowt.Scan(&tuteur.IdInterneMdt, &tuteur.Qualite, &tuteur.DatDbtMdt); err != nil {
// 					log.Fatalf("Error scanning tuteur curateur row: %v\n\n", err)
// 				}
// 				if tuteur.IdInterneMdt != "" {
// 					tuteurs = append(tuteurs, tuteur)
// 				}
// 			}
// 			if len(tuteurs) > 0 {
// 				pp.TuteurCurateur = &tuteurs
// 			}
// 		}
// 		//end tuteurs curateurs

// 		//cas des employeurs
// 		rowe, err := db.Query("SELECT IdInterneEmpl,DenominationSociale,RCCM,NIF,NIFP,DateCreation,DateEntree  FROM DIAMA.DBANK_EMPLOYEUR WHERE IdInterneClt = :1", pp.IdInterneClt)
// 		if err != nil {
// 			log.Fatalf("Error executing query for employeur: %v", err)
// 		}
// 		defer rowe.Close()
// 		var employeur models.PersonnePhysiqueEmployeur
// 		if rowe.Next() {
// 			if err := rowe.Scan(&employeur.IdInterneEmpl, &employeur.DenominationSociale, &employeur.RCCM, &employeur.NIF, &employeur.NIFP, &employeur.DateCreation, &employeur.DateEntree); err != nil {
// 				log.Fatalf("Error scanning employeur row: %v\n\n", err)
// 			}
// 			pp.Employeur = &employeur
// 		}
// 		//end employeurs

// 		//cas des donnees additionnelles
// 		rowa, err := db.Query("SELECT Cle,Valeur FROM DIAMA.DBANK_PP_ADDITIONNELLES WHERE IdInterneClt = :1", pp.IdInterneClt)
// 		if err != nil {
// 			log.Fatalf("Error executing query for donnees additionnelles: %v", err)
// 		}
// 		defer rowa.Close()
// 		var additionnelles models.PersonnePhysiqueAdditionnelles
// 		if rowa.Next() {
// 			if err := rowa.Scan(&additionnelles.Cle, &additionnelles.Valeur); err != nil {
// 				log.Fatalf("Error scanning donnees additionnelles row: %v\n\n", err)
// 			}
// 			pp.DonneesAdditionnelles = &additionnelles
// 		}
// 		//end donnees additionnelles

// 		arr_pp = append(arr_pp, pp)
// 	}

// 	//Préparation de la structure de réponse
// 	declaration := config.DeclarationPP{
// 		// NumDec:           strconv.Itoa(NumDeclarationPP),
// 		NumDec:           fmt.Sprintf("%04d", NumDeclarationPP),
// 		PartEmtr:         "038",
// 		TypDec:           "01",
// 		NbrDec:           strconv.Itoa(len(arr_pp)),
// 		DateDec:          strings.Join(strings.Split(time.Now().Format("02-01-2006"), "-"), ""),
// 		PersonnePhysique: arr_pp,
// 	}
// 	response := config.Response{
// 		Meta: config.Meta{
// 			Status:      200,
// 			Message:     "success",
// 			RequestTime: time.Now().Format("02-01-2006 15:04:05"),
// 			RequestID:   uuid.NewString(),
// 		},
// 		Data: config.Data{
// 			Declaration: declaration,
// 		},
// 	}

// 	w.Header().Set("Content-Type", "application/xml;charset='utf-8'")
// 	xmlData, err := xml.Marshal(response)
// 	if err != nil {
// 		http.Error(w, "Erreur lors de l'encodage XML", http.StatusInternalServerError)
// 		return
// 	}
// 	// w.Write([]byte(headerXml))
// 	w.Write([]byte(headerXml))
// 	w.Write(xmlData)

// 	arr_pp = nil
// }
