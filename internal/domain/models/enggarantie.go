package models

import (
	"database/sql"
)

// DBANK_GARANTIE
type EngagementsGarantie struct {
	Nooper         string       `xml:"-"`
	Numseq         int          `xml:"-"`
	Cliprt         string       `xml:"-"`
	EstDeclare     string       `xml:"-"`
	DateDeclare    sql.NullTime `xml:"-"`
	Datmaj         sql.NullTime `xml:"-"`
	RefIntGar      string       `xml:"RefIntGar,attr"`
	TypGar         string       `xml:"TypGar,attr"`
	DesGar         string       `xml:"DesGar,attr,omitempty"`
	CodDev         string       `xml:"CodDev,attr"`
	MntGar         float64      `xml:"MntGar,attr"`
	TypIdent       string       `xml:"TypIdent,attr,omitempty"`
	CodIdent       string       `xml:"CodIdent,attr,omitempty"`
	DatEval        string       `xml:"DatEval,attr,omitempty"`
	DatExp         string       `xml:"DatExp,attr,omitempty"`
	MntAffecGar    float64      `xml:"MntAffecGar,attr"`
	StatutGarantie string       `xml:"StatutGarantie,attr"`
	IdIntGarant    string       `xml:"IdIntGarant,attr,omitempty"`
}
