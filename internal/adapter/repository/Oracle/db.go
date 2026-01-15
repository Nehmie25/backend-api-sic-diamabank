package oracle

import (
	"database/sql"
	"log"

	_ "github.com/godror/godror"
)

func ConnectDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("godror", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}


	log.Println("Connexion DB établie")
	return db, nil
}
