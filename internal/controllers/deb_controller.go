package controllers

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"sicbcrg.diamabank.com/internal/models"
	"sicbcrg.diamabank.com/internal/services"
)

var NumDeclarationDEB int = 0

// DeclarationDEB defines the structure for Compte Debiteurs declaration
type DeclarationDEB struct {
	NumDec          string                   `xml:"NumDec,attr"`
	PartEmtr        string                   `xml:"PartEmtr,attr"`
	TypDec          string                   `xml:"TypDec,attr"`
	NbrDec          string                   `xml:"NbrDec,attr"`
	DateDec         string                   `xml:"DateDec,attr"`
	DatArr          string                   `xml:"DatArr,attr"`
	CompteDebiteurs []models.CompteDebiteurs `xml:"CompteDebiteur"`
}

// HandlerCompteDebiteurs handles Compte Debiteurs requests using service layer
func HandlerCompteDebiteurs(service *services.CompteDebiteurService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dateDeclaration string = r.URL.Query().Get("date")
		if dateDeclaration == "" {
			http.Error(w, "Saisir une date valide", http.StatusBadRequest)
			return
		}

		ctx := r.Context()

		// Exécuter la procédure stockée via le service
		if err := service.ExecuteDeclaration(ctx, dateDeclaration); err != nil {
			log.Printf("Error executing declaration: %v", err)
			http.Error(w, "Erreur lors de l'exécution de la procédure", http.StatusInternalServerError)
			return
		}

		fmt.Printf("Date de la déclaration : %s\n", dateDeclaration)

		// Récupérer tous les comptes débiteurs avec leurs détails via le service
		arr_deb, err := service.GetAllComptesWithDetails(ctx)
		if err != nil {
			log.Printf("Error fetching comptes debiteurs: %v", err)
			http.Error(w, "Erreur lors de la récupération des données", http.StatusInternalServerError)
			return
		}

		NumDeclarationDEB++

		// Parse date for DatArr attribute
		date, err := time.Parse("02/01/06", dateDeclaration)
		if err != nil {
			log.Printf("Error parsing date: %v", err)
			http.Error(w, "Format de date invalide", http.StatusBadRequest)
			return
		}

		// Préparation de la structure de réponse
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
	}
}
