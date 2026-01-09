package models 

import (
	"database/sql"
)

// DBANK_TUTEURCURATEUR
type PersonnePhysiqueTuteurCurateur struct {
	Client       string       `xml:"-"`
	IDp          string       `xml:"-"`
	EstDeclare   string       `xml:"-"`
	DateDeclare  sql.NullTime `xml:"-"`
	Datmaj       sql.NullTime `xml:"-"`
	IdInterneClt string       `xml:"-"`
	IdInterneMdt string       `xml:"IdInterneMdt,attr"`
	Qualite      string       `xml:"Qualite,attr"`
	DatDbtMdt    string       `xml:"DatDbtMdt,attr"`
}
