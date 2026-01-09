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

var NumDeclarationPP int = 0

var headerXml string = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>` + "\n"

// DeclarationPP defines the structure for Personne Physique declaration
type DeclarationPP struct {
	NumDec           string                    `xml:"NumDec,attr"`
	PartEmtr         string                    `xml:"PartEmtr,attr"`
	TypDec           string                    `xml:"TypDec,attr"`
	NbrDec           string                    `xml:"NbrDec,attr"`
	DateDec          string                    `xml:"DateDec,attr"`
	PersonnePhysique []models.PersonnePhysique `xml:"PersonnePhysique"`
}

type Meta struct {
	Status      int    `xml:"status"`
	Message     string `xml:"message"`
	RequestTime string `xml:"requestTime"`
	RequestID   string `xml:"requestId"`
}

type Data struct {
	Declaration interface{} `xml:"declaration"`
}

type Response struct {
	Meta Meta `xml:"meta"`
	Data Data `xml:"data"`
}

// HandlerPersonnePhysique handles Personne Physique requests using service layer
func HandlerPersonnePhysique(service *services.PersonnePhysiqueService) http.HandlerFunc {
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

		// Récupérer toutes les personnes physiques avec leurs détails via le service
		arr_pp, err := service.GetAllPersonnesWithDetails(ctx)
		if err != nil {
			log.Printf("Error fetching personnes: %v", err)
			http.Error(w, "Erreur lors de la récupération des données", http.StatusInternalServerError)
			return
		}

		NumDeclarationPP++

		// Préparation de la structure de réponse
		declaration := DeclarationPP{
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

		w.Write([]byte(headerXml))
		w.Write(xmlData)
	}
}
