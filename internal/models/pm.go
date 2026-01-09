package models

import (
	"database/sql"
)

// DBANK_PERSMORALE
type PersonneMorale struct {
	Client           string       `xml:"-"`
	EstDeclare       string       `xml:"-"`
	DateDeclare      sql.NullTime `xml:"-"`
	Datmaj           sql.NullTime `xml:"-"`
	NatDec           string       `xml:"NatDec,attr"`
	NatClient        string       `xml:"NatClient,attr"`
	IdInterneClt     string       `xml:"IdInterneClt,attr"`
	DenomSocial      string       `xml:"DenomSocial,attr"`
	Sigle            string       `xml:"Sigle,attr,omitempty"`
	DatCreat         string       `xml:"DatCreat,attr"`
	Statut           string       `xml:"Statut,attr"`
	DatCreaPart      string       `xml:"DatCreaPart,attr"`
	FormeJuridique   string       `xml:"FormeJuridique,attr"`
	PaysSiegeSocial  string       `xml:"PaysSiegeSocial,attr"`
	VilleSiegeSocial string       `xml:"VilleSiegeSocial,attr"`
	Mobile           string       `xml:"Mobile,attr"`
	Email            string       `xml:"Email,attr,omitempty"`
	SiteWeb          string       `xml:"SiteWeb,attr,omitempty"`
	Adress           string       `xml:"Adress,attr"`
	CommuneAdress    string       `xml:"CommuneAdress,attr,omitempty"`
	CodePostal       string       `xml:"CodePostal,attr,omitempty"`
	Resident         string       `xml:"Resident,attr"`
	RCCM             string       `xml:"RCCM,attr,omitempty"`
	NIF              string       `xml:"NIF,attr,omitempty"`
	NIFP             string       `xml:"NIFP,attr,omitempty"`
	NumAgrement      string       `xml:"NumAgrement,attr,omitempty"`
	NumSecSoc        string       `xml:"NumSecSoc,attr,omitempty"`
	SecActEcon       string       `xml:"ActEcon,attr"`
	SectInst         string       `xml:"SectInst,attr"`
	SitBancaire      string       `xml:"SitBancaire,attr,omitempty"`
	DateDebIB        string       `xml:"DateDebIB,attr,omitempty"`
	DateFinIB        string       `xml:"DateFinIB,attr,omitempty"`

	Mandataires           *[]PersonneMoraleMandataire     `xml:"Mandataire,omitempty"`
	CompteAssocie         *[]PersonneMoraleCompteAssocie  `xml:"CompteAssocie,omitempty"`
	Actionnaire           *[]PersonneMoraleActionnaire    `xml:"Actionnaire,omitempty"`
	DonneesAdditionnelles *[]PersonneMoraleAdditionnelles `xml:"DonneesAdditionnelles,omitempty"`
}
