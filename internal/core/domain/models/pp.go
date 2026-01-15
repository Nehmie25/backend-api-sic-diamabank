package models

import (
	"database/sql"
)

// DBANK_PERSPHYSIQUE
type PersonnePhysique struct {
	Client        string       `xml:"-"`
	Idp           string       `xml:"-"`
	EstDeclare    string       `xml:"-"`
	DateDeclare   sql.NullTime `xml:"-"`
	Datmaj        sql.NullTime `xml:"-"`
	NatDec        string       `xml:"NatDec,attr"`
	NatClient     string       `xml:"NatClient,attr"`
	NIN           string       `xml:"NIN,attr,omitempty"`
	IdInterneClt  string       `xml:"IdInterneClt,attr"`
	DatCreaPart   string       `xml:"DatCreaPart,attr"`
	NomNaiClt     string       `xml:"NomNaiClt,attr"`
	NomMtlClt     string       `xml:"NomMtlClt,attr,omitempty"`
	PrenomClt     string       `xml:"PrenomClt,attr"`
	Sexe          string       `xml:"Sexe,attr"`
	DatNai        string       `xml:"DatNai,attr"`
	EtatCivil     string       `xml:"EtatCivil,attr"`
	NomPere       string       `xml:"NomPere,attr"`
	PrenomPere    string       `xml:"PrenomPere,attr"`
	NomNaiMere    string       `xml:"NomNaiMere,attr"`
	PrmMre        string       `xml:"PrmMre,attr"`
	VilleNai      string       `xml:"VilleNai,attr"`
	PaysNai       string       `xml:"PaysNai,attr"`
	NatClt        string       `xml:"NatClt,attr"`
	Resident      string       `xml:"Resident,attr"`
	PaysRes       string       `xml:"PaysRes,attr"`
	Mobile        string       `xml:"Mobile,attr"`
	Email         string       `xml:"Email,attr,omitempty"`
	Adress        string       `xml:"Adress,attr"`
	CommuneAdress string       `xml:"CommuneAdress,attr,omitempty"`
	CodePostal    string       `xml:"CodePostal,attr,omitempty"`
	Profession    string       `xml:"Profession,attr,omitempty"`
	SecActEcon    string       `xml:"ActEcon,attr,omitempty"`
	SectInst      string       `xml:"SectInst,attr"`
	NumSecSoc     string       `xml:"NumSecSoc,attr,omitempty"`
	STutelle      string       `xml:"STutelle,attr"`
	StatutClt     string       `xml:"StatutClt,attr"`
	DateDeces     string       `xml:"DateDeces,attr,omitempty"`
	SitBancaire   string       `xml:"SitBancaire,attr,omitempty"`
	DateDebIB     string       `xml:"DateDebIB,attr,omitempty"`
	DateFinIB     string       `xml:"DateFinIB,attr,omitempty"`

	CompteAssocie         *[]PersonnePhysiqueCompteAssocie  `xml:"CompteAssocie,omitempty"`
	Piece                 *[]PersonnePhysiquePiece          `xml:"Piece,omitempty"`
	DonneeComplementaire  *PersonnePhysiqueComplementaire   `xml:"DonneeComplementaire,omitempty"`
	TuteurCurateur        *[]PersonnePhysiqueTuteurCurateur `xml:"TuteurCurateur,omitempty"`
	Employeur             *PersonnePhysiqueEmployeur        `xml:"Employeur,omitempty"`
	DonneesAdditionnelles *PersonnePhysiqueAdditionnelles   `xml:"DonneesAdditionnelles,omitempty"`
}
