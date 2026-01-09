package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/godror/godror"
	"sicbcrg.diamabank.com/internal/database"
	"sicbcrg.diamabank.com/internal/router"
	"sicbcrg.diamabank.com/internal/repositories"
	"sicbcrg.diamabank.com/internal/services"
)



func main() {
	// Initialize database
	err := database.Init()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.DB.Close()


		// Créer les repositories
		ppRepo := repositories.NewOraclePersonnePhysiqueRepository(database.DB)
		pmRepo := repositories.NewOraclePersonneMoraleRepository(database.DB)
		engRepo := repositories.NewOracleEngagementsRepository(database.DB)
		encRepo := repositories.NewOracleEncoursRepository(database.DB)
		debRepo := repositories.NewOracleCompteDebiteurRepository(database.DB)

		// Créer les services
		ppService := services.NewPersonnePhysiqueService(ppRepo)
		pmService := services.NewPersonneMoraleService(pmRepo)
		engService := services.NewEngagementsService(engRepo)
		encService := services.NewEncoursService(encRepo)
		debService := services.NewCompteDebiteurService(debRepo)

		// Injecter les services dans le router
		apiRouter := router.SetupRoutes(ppService, pmService, engService, encService, debService)

	// Configuration serveur depuis variables d'environnement
	apiPort := os.Getenv("API_PORT")
	if apiPort == "" {
		apiPort = "8080"
	}
	serverAddr := "0.0.0.0:" + apiPort

	server := http.Server{
		Addr:         serverAddr,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  10 * time.Second,
		Handler:      apiRouter,
	}

	fmt.Printf("API Server listening on %s\n", serverAddr)
	log.Fatal(server.ListenAndServe())
}