package models

import (
	"database/sql"
)

// DIAMA.DBANK_PERSMORALE_CPT
type PersonneMoraleCompteAssocie struct {
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
	StatCpt      string       `xml:"StatCpt,attr"`

	MandataireAssocie *[]PersonneMoraleMandataireMandataireAssocie `xml:"MandataireAssocie,omitempty"`
}
