package models


type Historique struct {
	ID        	int       	`json:"id"`
	Operation 	string    	`json:"operation"`
	UserID    	string    	`json:"user_id"`
	UserNames  	string    	`json:"user_name"`
}
