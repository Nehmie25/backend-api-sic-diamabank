package models

import(
	"time"
)

// DBANK_EMPLOYEUR
type User struct {
	Id           string     	`json:"id"`
	Nom          string     	`json:"nom"`
	Email        string 		`json:"email"`
	MotDePasse   string  		`json:"motDePasse"`
	Role         string  		`json:"role"`
	Isactive		bool   		`json:"isactive"`
	CreatedAt    time.Time  	`json:"createdAt"`
	UpdatedAt    time.Time  	`json:"updatedAt"`
}