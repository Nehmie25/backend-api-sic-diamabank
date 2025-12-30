package models

import (
	"database/sql"
)

// DBANK_MANDATAIRE
type PersonneMoraleMandataire struct {
	Client       string       `xml:"-"`
	Idp          string       `xml:"-"`
	EstDeclare   string       `xml:"-"`
	DateDeclare  sql.NullTime `xml:"-"`
	Datmaj       sql.NullTime `xml:"-"`
	IdInterneClt string       `xml:"-"`
	IdInterneMdt string       `xml:"IdInterneMdt,attr"`
	Qualite      string       `xml:"Qualite,attr"`
	DatDebMdt    string       `xml:"DatDebMdt,attr,omitempty"`
	DatFinMdt    string       `xml:"DatFinMdt,attr,omitempty"`
}
