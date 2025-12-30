package models

import (
	"database/sql"
)

// DBANK_PP_ADDITIONNELLES
type PersonnePhysiqueAdditionnelles struct {
	Client       string       `xml:"-"`
	EstDeclare   string       `xml:"-"`
	DateDeclare  sql.NullTime `xml:"-"`
	Datmaj       sql.NullTime `xml:"-"`
	IdInterneClt string       `xml:"-"`
	Cle          string       `xml:"Cle,attr"`
	Valeur       string       `xml:"Valeur,attr"`
}
