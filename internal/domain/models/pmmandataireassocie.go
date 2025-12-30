package models

import (
	"database/sql"
)

// DBANK_MANDASSOCIE
type PersonneMoraleMandataireMandataireAssocie struct {
	Client          string       `xml:"-"`
	Idp             string       `xml:"-"`
	EstDeclare      string       `xml:"-"`
	DateDeclare     sql.NullTime `xml:"-"`
	Datmaj          sql.NullTime `xml:"-"`
	IdInterneClt    string       `xml:"-"`
	IdInterneMdtCpt string       `xml:"IdInterneMdtCpt,attr"`
	DatDebMdt       string       `xml:"DatDebMdt,attr,omitempty"`
	DatFinMdt       string       `xml:"DatFinMdt,attr,omitempty"`
}
