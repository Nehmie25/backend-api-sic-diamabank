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

var NumDeclarationENC int = 0

// DeclarationENC defines the structure for Encours declaration
type DeclarationENC struct {
	NumDec            string           `xml:"NumDec,attr"`
	PartEmtr          string           `xml:"PartEmtr,attr"`
	TypDec            string           `xml:"TypDec,attr"`
	NbrDec            string           `xml:"NbrDec,attr"`
	DateDec           string           `xml:"DateDec,attr"`
	DatArr            string           `xml:"DatArr,attr"`
	EncoursEngagement []models.Encours `xml:"EncoursEngagement"`
}

// HandlerEncours handles Encours requests using service layer
func HandlerEncours(service *services.EncoursService) http.HandlerFunc {
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

		// Récupérer tous les encours avec leurs détails via le service
		arr_enc, err := service.GetAllEncoursWithDetails(ctx)
		if err != nil {
			log.Printf("Error fetching encours: %v", err)
			http.Error(w, "Erreur lors de la récupération des données", http.StatusInternalServerError)
			return
		}

		NumDeclarationENC++

		// Parse date for DatArr attribute
		date, err := time.Parse("02/01/06", dateDeclaration)
		if err != nil {
			log.Printf("Error parsing date: %v", err)
			http.Error(w, "Format de date invalide", http.StatusBadRequest)
			return
		}

		// Préparation de la structure de réponse
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
	}
}
