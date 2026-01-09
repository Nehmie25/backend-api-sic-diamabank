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

var NumDeclarationPM int = 0

// DeclarationPM defines the structure for Personne Morale declaration
type DeclarationPM struct {
	NumDec         string                  `xml:"NumDec,attr"`
	PartEmtr       string                  `xml:"PartEmtr,attr"`
	TypDec         string                  `xml:"TypDec,attr"`
	NbrDec         string                  `xml:"NbrDec,attr"`
	DateDec        string                  `xml:"DateDec,attr"`
	PersonneMorale []models.PersonneMorale `xml:"PersonneMorale"`
}

// HandlerPersonneMorale handles Personne Morale requests using service layer
func HandlerPersonneMorale(service *services.PersonneMoraleService) http.HandlerFunc {
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

		// Récupérer toutes les personnes morales avec leurs détails via le service
		arr_pm, err := service.GetAllPersonnesWithDetails(ctx)
		if err != nil {
			log.Printf("Error fetching personnes morales: %v", err)
			http.Error(w, "Erreur lors de la récupération des données", http.StatusInternalServerError)
			return
		}

		NumDeclarationPM++

		// Préparation de la structure de réponse
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
		xmlData, err := xml.Marshal(response)
		if err != nil {
			http.Error(w, "Erreur lors de l'encodage XML", http.StatusInternalServerError)
			return
		}

		w.Write([]byte(headerXml))
		w.Write(xmlData)
	}
}
