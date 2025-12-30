package models

import (
	"database/sql"
)

// DBANK_PP_PIECE
type PersonnePhysiquePiece struct {
	Client       string       `xml:"-"`
	EstDeclare   string       `xml:"-"`
	DateDeclare  sql.NullTime `xml:"-"`
	Datmaj       sql.NullTime `xml:"-"`
	IdInterneClt string       `xml:"-"`
	TypPiece     string       `xml:"TypPiece,attr"`
	NumPiece     string       `xml:"NumPiece,attr"`
	DatEmiPiece  string       `xml:"DatEmiPiece,attr"`
	LieuEmiPiece string       `xml:"LieuEmiPiece,attr"`
	PaysEmiPiece string       `xml:"PaysEmiPiece,attr"`
	FinValPiece  string       `xml:"FinValPiece,attr"`
}
