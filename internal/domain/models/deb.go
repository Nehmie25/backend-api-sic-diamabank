package models

import (
	"database/sql"
)

// DBANK_CPTDEBITEURS
type CompteDebiteurs struct {
	Nooper        string       `xml:"-"`
	Numseq        int          `xml:"-"`
	Datrmb        sql.NullTime `xml:"-"`
	Cliprt        string       `xml:"-"`
	Cptprt        string       `xml:"-"`
	Cptvue        string       `xml:"-"`
	Cptimp        string       `xml:"-"`
	EstDeclare    string       `xml:"-"`
	DateDeclare   sql.NullTime `xml:"-"`
	Datmaj        sql.NullTime `xml:"-"`
	NatDec        string       `xml:"NatDec,attr"`
	CodDev        string       `xml:"CodDev,attr"`
	Rib           string       `xml:"Rib,attr"`
	SoldeDeb      string       `xml:"SoldeDeb,attr"`
	DateDefaill   string       `xml:"DateDefaill,attr,omitempty"`
	NbrJourDebMax string       `xml:"NbrJourDebMax,attr"`
	SoldeDebMax   string       `xml:"SoldeDebMax,attr"`
	MntProv       string       `xml:"MntProv,attr"`
	MntPerte      string       `xml:"MntPerte,attr"`
	MntAgi        string       `xml:"MntAgi,attr,omitempty"`
	QualiCre      string       `xml:"QualiCre,attr"`
	IdIntTit      string       `xml:"IdIntTit,attr"`

	Titulaire             CompteDebiteursTitulaire         `xml:"Titulaire,omitempty"`
	DonneesAdditionnelles *[]CompteDebiteursAdditionnelles `xml:"DonneesAdditionnelles,omitempty"`
}
