package models

import (
	"database/sql"
)

// DBANK_PM_ADDITIONNELLES
type PersonneMoraleAdditionnelles struct {
	Client       string       `xml:"-"`
	EstDeclare   string       `xml:"-"`
	DateDeclare  sql.NullTime `xml:"-"`
	Datmaj       sql.NullTime `xml:"-"`
	IdInterneClt string       `xml:"-"`
	Cle          string       `xml:"Cle,attr"`
	Valeur       string       `xml:"Valeur,attr"`
}
