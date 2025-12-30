package models

import (
	"database/sql"
)

// DBANK_PP_COMPLEMENTAIRE
type PersonnePhysiqueComplementaire struct {
	Client       string       `xml:"-"`
	EstDeclare   string       `xml:"-"`
	DateDeclare  sql.NullTime `xml:"-"`
	Datmaj       sql.NullTime `xml:"-"`
	IdInterneClt string       `xml:"-"`
	NbPersCharge string       `xml:"NbPersCharge,attr,omitempty"`
	RevMensMoy   string       `xml:"RevMensMoy,attr,omitempty"`
	DepMensMoy   string       `xml:"DepMensMoy,attr,omitempty"`
	PropLoc      string       `xml:"PropLoc,attr,omitempty"`
}
