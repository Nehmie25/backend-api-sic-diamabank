package models

import (
	"database/sql"
)

// DBANK_ENC_ADDITIONNELLES
type EncoursAdditionnelles struct {
	Nooper      string       `xml:"-"`
	Numseq      int          `xml:"-"`
	Cliprt      string       `xml:"-"`
	EstDeclare  string       `xml:"-"`
	DateDeclare sql.NullTime `xml:"-"`
	Datmaj      sql.NullTime `xml:"-"`
	RefIntEng   string       `xml:"-"`
	Cle         string       `xml:"Cle,attr"`
	Valeur      string       `xml:"Valeur,attr"`
}
