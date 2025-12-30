package models

import (
	"database/sql"
)

// DBANK_EMPLOYEUR
type PersonnePhysiqueEmployeur struct {
	Client              string       `xml:"-"`
	EstDeclare          string       `xml:"-"`
	DateDeclare         sql.NullTime `xml:"-"`
	Datmaj              sql.NullTime `xml:"-"`
	IdInterneClt        string       `xml:"-"`
	IdInterneEmpl       string       `xml:"IdInterneEmpl,attr,omitempty"`
	DenominationSociale string       `xml:"DenominationSociale,attr,omitempty"`
	RCCM                string       `xml:"RCCM,attr,omitempty"`
	NIF                 string       `xml:"NIF,attr,omitempty"`
	NIFP                string       `xml:"NIFP,attr,omitempty"`
	DateCreation        string       `xml:"DateCreation,attr,omitempty"`
	DateEntree          string       `xml:"DateEntree,attr"`
}
