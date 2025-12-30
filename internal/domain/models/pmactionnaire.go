package models

import (
	"database/sql"
)

// DBANK_ACTIONNAIRE
type PersonneMoraleActionnaire struct {
	Client       string       `xml:"-"`
	Idp          string       `xml:"-"`
	EstDeclare   string       `xml:"-"`
	DateDeclare  sql.NullTime `xml:"-"`
	Datmaj       sql.NullTime `xml:"-"`
	IdInterneClt string       `xml:"-"`
	IdInterneAct string       `xml:"IdInterneAct,attr"`
	PartAct      string       `xml:"PartAct,attr"`
	DaEntrAct    string       `xml:"DaEntrAct,attr,omitempty"`
}
