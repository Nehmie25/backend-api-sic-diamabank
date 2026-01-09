package models

import (
	"database/sql"
)

// DBANK_ENCOURS
type Encours struct {
	Nooper      string       `xml:"-"`
	Numseq      string       `xml:"-"`
	Datrmb      sql.NullTime `xml:"-"`
	Datrmbant   sql.NullTime `xml:"-"`
	Cliprt      string       `xml:"-"`
	Cptprt      string       `xml:"-"`
	Cptvue      string       `xml:"-"`
	Cptimp      string       `xml:"-"`
	EstDeclare  string       `xml:"-"`
	DateDeclare sql.NullTime `xml:"-"`
	Datmaj      sql.NullTime `xml:"-"`
	NatDec      string       `xml:"NatDec,attr"`
	RefIntEng   string       `xml:"RefIntEng,attr"`
	CodDev      string       `xml:"CodDev,attr"`
	DatEch      string       `xml:"DatEch,attr"`
	MntDerEch   string       `xml:"MntDerEch,attr"`
	MonPai      string       `xml:"MonPai,attr"`
	DatPai      string       `xml:"DatPai,attr"`
	MntHBil     string       `xml:"MntHBil,attr"`
	MntRemAnt   string       `xml:"MntRemAnt,attr"`
	MntCRDU     string       `xml:"MntCRDU,attr"`
	MntCreRat   string       `xml:"MntCreRat,attr"`
	MntUtilise  string       `xml:"MntUtilise,attr"`
	MntAgi      string       `xml:"MntAgi,attr,omitempty"`
	MntCapImp   string       `xml:"MntCapImp,attr"`
	MntTotImp   string       `xml:"MntTotImp,attr"`
	DatDefaill  string       `xml:"DatDefaill,attr,omitempty"`
	MntPro      string       `xml:"MntPro,attr"`
	MntPerte    string       `xml:"MntPerte,attr"`
	NbrEchPay   string       `xml:"NbrEchPay,attr"`
	NbrEchImp   string       `xml:"NbrEchImp,attr"`
	NbrEchRest  string       `xml:"NbrEchRest,attr"`
	QualiCre    string       `xml:"QualiCre,attr"`
	PD          string       `xml:"PD,attr,omitempty"`
	LGD         string       `xml:"LGD,attr"`
	CCF         string       `xml:"CCF,attr,omitempty"`
	IFRSStage   string       `xml:"IFRSStage,attr,omitempty"`
	DatEvent    string       `xml:"DatEvent,attr"`

	DonneesAdditionnelles *[]EncoursAdditionnelles `xml:"DonneesAdditionnelles,omitempty"`
}
