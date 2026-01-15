package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"time"

	//"github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() (db *sql.DB) {
	portStr := os.Getenv("DB_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatal("Invalid DB_PORT:", err)
	}

	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		url.QueryEscape(os.Getenv("DB_PASSWORD")),
		os.Getenv("DB_HOST"),
		port,
		os.Getenv("DB_NAME"))

	DB, err = sql.Open("postgres", connString)
	if err != nil {
		log.Fatal("Connexion PostgreSQL échouée:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := DB.PingContext(ctx); err != nil {
		log.Fatal("PostgreSQL inaccessible:", err)
	}

	log.Println("PostgreSQL connecté avec succès")

	return DB
}
