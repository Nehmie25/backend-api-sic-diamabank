package models

import (
	"database/sql"
)

// DBANK_ENGAGEMENTS
type Engagements struct {
	Nooper           string       `xml:"-"`
	Numseq           int          `xml:"-"`
	EstDeclare       string       `xml:"-"`
	DateDeclare      sql.NullTime `xml:"-"`
	Datmaj           sql.NullTime `xml:"-"`
	NatDec           string       `xml:"NatDec,attr"`
	TypEve           string       `xml:"TypEve,attr"`
	RefIntEng        string       `xml:"RefIntEng,attr,omitempty"`
	LigneParent      string       `xml:"LigneParent,attr"`
	RefIntLigne      string       `xml:"RefIntLigne,attr,omitempty"`
	RefDemandeEng    string       `xml:"RefDemandeEng,attr,omitempty"`
	DatDem           string       `xml:"DatDem,attr,omitempty"`
	TypeModif        string       `xml:"TypeModif,attr"`
	EstDout          string       `xml:"EstDout,attr,omitempty"`
	Cloture          string       `xml:"Cloture,attr"`
	DatAccord        string       `xml:"DatAccord,attr,omitempty"`
	MotifCloture     string       `xml:"MotifCloture,attr,omitempty"`
	DatClo           string       `xml:"DatClo,attr,omitempty"`
	DateMEP          string       `xml:"DateMEP,attr"`
	TypEng           string       `xml:"TypEng,attr"`
	MntEng           string       `xml:"MntEng,attr"`
	MntInt           string       `xml:"MntInt,attr,omitempty"`
	CodDev           string       `xml:"CodDev,attr"`
	PeriodRemb       string       `xml:"PeriodRemb,attr"`
	TxIntEng         string       `xml:"TxIntEng,attr,omitempty"`
	TypTxInt         string       `xml:"TypTxInt,attr,omitempty"`
	TxComm           string       `xml:"TxComm,attr,omitempty"`
	IndRef           string       `xml:"IndRef,attr,omitempty"`
	Sprd             string       `xml:"Sprd,attr,omitempty"`
	TxEffGlob        string       `xml:"TxEffGlob,attr,omitempty"`
	MoyRemb          string       `xml:"MoyRemb,attr,omitempty"`
	TypAmo           string       `xml:"TypAmo,attr,omitempty"`
	TypDiffAmo       string       `xml:"TypDiffAmo,attr,omitempty"`
	UnitDur          string       `xml:"UnitDur,attr,omitempty"`
	PerDiffAmo       string       `xml:"PerDiffAmo,attr,omitempty"`
	MntEch           string       `xml:"MntEch,attr,omitempty"`
	NbrEch           string       `xml:"NbrEch,attr,omitempty"`
	DatPremEch       string       `xml:"DatPremEch,attr,omitempty"`
	DatFin           string       `xml:"DatFin,attr"`
	MntFrais         string       `xml:"MntFrais,attr"`
	MntComm          string       `xml:"MntComm,attr"`
	CodAgce          string       `xml:"CodAgce,attr"`
	EstRachatCreance string       `xml:"EstRachatCreance,attr"`
	ParCont          string       `xml:"ParCont,attr,omitempty"`
	ValNom           string       `xml:"ValNom,attr,omitempty"`
	ValCess          string       `xml:"ValCess,attr,omitempty"`
	DatEvent         string       `xml:"DatEvent,attr"`
	IdIntBen         string       `xml:"-"`
	PourBenef        string       `xml:"-"`

	Beneficiaire          *[]EngagementsBeneficiaire   `xml:"Beneficiaire,omitempty"`
	Garantie              *[]EngagementsGarantie       `xml:"Garantie,omitempty"`
	Consolidation         *[]EngagementsConsolidation  `xml:"Consolidation,omitempty"`
	DonneesAdditionnelles *[]EngagementsAdditionnelles `xml:"DonneesAdditionnelles,omitempty"`
}
