package models

// DBANK_ENGAGEMENTS
type EngagementsBeneficiaire struct {
	IdIntBen  string  `xml:"IdIntBen,attr"`
	PourBenef float64 `xml:"PourBenef,attr"`
}
