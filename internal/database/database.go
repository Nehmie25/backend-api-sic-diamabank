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
		log.Println("DB_USER not set in environment")
	}
	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		log.Println("DB_PASSWORD not set in environment")
	}
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		log.Println("DB_HOST not set in environment")
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		log.Println("DB_PORT not set in environment")
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		log.Println("DB_NAME not set in environment")
	}

	dataSourceName := fmt.Sprintf(`user="%s" password="%s" connectString="%s:%s/%s"`, dbUser, dbPassword, dbHost, dbPort, dbName)
	
	var dbErr error
	DB, dbErr = sql.Open("godror", dataSourceName)
	if dbErr != nil {
		return fmt.Errorf("error opening database connection: %v", dbErr)
	}

	dbErr = DB.Ping()
	if dbErr != nil {
		return fmt.Errorf("error connecting to the database: %v", dbErr)
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
