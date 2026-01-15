package config

import (
	"GoWebapitest/internal/core/domain/models"
)

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

type JsonData struct {
	Users interface{} `xml:"declaration"`
}

type Meta struct {
	Status      int    `xml:"status"`
	Message     string `xml:"message"`
	RequestTime string `xml:"requestTime"`
	RequestID   string `xml:"requestId"`
}


type JsonMeta struct {
	Status      int    `xml:"status"`
	Message     string `xml:"message"`
	RequestTime string `xml:"requestTime"`
	RequestID   string `xml:"requestId"`
}

type Response struct {
	Meta Meta `xml:"meta"`
	Data Data `xml:"data"`
}

type JsonResponse struct {
	Meta JsonMeta `json:"meta"`
	Data JsonData `json:"data"`
}
