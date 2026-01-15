package models

import (
	"database/sql"
)

// DBANK_PERSPHYSIQUE_CPT
type PersonnePhysiqueCompteAssocie struct {
	Client       string       `xml:"-"`
	Compte       string       `xml:"-"`
	Devise       string       `xml:"-"`
	Ncg          string       `xml:"-"`
	EstDeclare   string       `xml:"-"`
	DateDeclare  sql.NullTime `xml:"-"`
	Datmaj       sql.NullTime `xml:"-"`
	IdInterneClt string       `xml:"-"`
	CodAgce      string       `xml:"CodAgce,attr"`
	NumCpt       string       `xml:"NumCpt,attr"`
	CleRib       string       `xml:"CleRib,attr,omitempty"`
	TypCpt       string       `xml:"TypCpt,attr"`
	StatCpt      string       `xml:"StatCpt,attr"`
}
