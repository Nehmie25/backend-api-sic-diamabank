package models

import (
	"database/sql"
)

// DBANK_MANDASSOCIE
type EngagementsConsolidation struct {
	Client          string       `xml:"Client,attr"`
	Idp             string       `xml:"Idp,attr"`
	EstDeclare      string       `xml:"EstDeclare,attr"`
	DateDeclare     sql.NullTime `xml:"DateDeclare,attr"`
	Datmaj          sql.NullTime `xml:"Datmaj,attr"`
	IdInterneClt    string       `xml:"IdInterneClt,attr"`
	IdInterneMdtCpt string       `xml:"IdInterneMdtCpt,attr"`
	DatDebMdt       string       `xml:"DatDebMdt,attr"`
	DatFinMdt       string       `xml:"DatFinMdt,attr"`
}
