package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/godror/godror"
)

var DB *sql.DB

func Init() error {
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "SICPROD"
	}
	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "NMu6F0DtoXFnnkyHt80Z"
	}
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "10.0.16.3"
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "1521"
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "ORCLPDB"
	}

	dataSourceName := fmt.Sprintf(`user="%s" password="%s" connectString="%s:%s/%s"`, dbUser, dbPassword, dbHost, dbPort, dbName)
	
	var err error
	DB, err = sql.Open("godror", dataSourceName)
	if err != nil {
		return fmt.Errorf("error opening database connection: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("error connecting to the database: %v", err)
	}
	
	log.Println("Successfully connected to Oracle Database!")
	return nil
}

func Close() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}
