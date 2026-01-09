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

var NumDeclarationENG int = 0

// DeclarationENG defines the structure for Engagements declaration
type DeclarationENG struct {
	NumDec     string               `xml:"NumDec,attr"`
	PartEmtr   string               `xml:"PartEmtr,attr"`
	TypDec     string               `xml:"TypDec,attr"`
	NbrDec     string               `xml:"NbrDec,attr"`
	DateDec    string               `xml:"DateDec,attr"`
	Engagement []models.Engagements `xml:"Engagement"`
}

// HandlerEngagement handles Engagements requests using service layer
func HandlerEngagement(service *services.EngagementsService) http.HandlerFunc {
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

		// Récupérer tous les engagements avec leurs détails via le service
		arr_eng, err := service.GetAllEngagementsWithDetails(ctx)
		if err != nil {
			log.Printf("Error fetching engagements: %v", err)
			http.Error(w, "Erreur lors de la récupération des données", http.StatusInternalServerError)
			return
		}

		NumDeclarationENG++

		// Préparation de la structure de réponse
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
	}
}
