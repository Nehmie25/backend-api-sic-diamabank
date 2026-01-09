package router

import (
	"encoding/json"
	"net/http"

	"sicbcrg.diamabank.com/internal/controllers"
	"sicbcrg.diamabank.com/internal/services"
)

// WelcomeResponse defines the welcome message structure
type WelcomeResponse struct {
	Message   string   `json:"message"`
	Status    string   `json:"status"`
	Version   string   `json:"version"`
	Endpoints []string `json:"endpoints"`
}

// homeHandler returns a welcome JSON response
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	
	response := WelcomeResponse{
		Message: "Bienvenue sur l'API SIC DiamBank",
		Status:  "En ligne",
		Version: "1.0.0",
		Endpoints: []string{
			"GET /declaration/personnephysique?date=DD/MM/YY",
			"GET /declaration/personnemorale?date=DD/MM/YY",
			"GET /declaration/engagements?date=DD/MM/YY",
			"GET /declaration/encours?date=DD/MM/YY",
			"GET /declaration/comptedebiteurs?date=DD/MM/YY",
		},
	}
	
	json.NewEncoder(w).Encode(response)
}

// SetupRoutes configures all API routes
func SetupRoutes(
	ppService *services.PersonnePhysiqueService,
	pmService *services.PersonneMoraleService,
	engService *services.EngagementsService,
	encService *services.EncoursService,
	debService *services.CompteDebiteurService,
) *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/declaration/personnephysique", controllers.HandlerPersonnePhysique(ppService))
	router.HandleFunc("/declaration/personnemorale", controllers.HandlerPersonneMorale(pmService))
	router.HandleFunc("/declaration/engagements", controllers.HandlerEngagement(engService))
	router.HandleFunc("/declaration/encours", controllers.HandlerEncours(encService))
	router.HandleFunc("/declaration/comptedebiteurs", controllers.HandlerCompteDebiteurs(debService))

	return router
}
