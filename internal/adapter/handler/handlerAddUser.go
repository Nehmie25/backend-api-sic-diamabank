package handler

import (
	"GoWebapitest/config"
	"GoWebapitest/internal/core/domain/models"
	"GoWebapitest/internal/core/service"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func HandlerAddUser(service *service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Récupérer les paramètres du formulaire par POST

		var body models.User

		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if body.Nom == "" || body.Email == "" || body.MotDePasse == "" || body.Role == "" {
			http.Error(w, "Missing required fields", http.StatusBadRequest)
			return
		}
		
		hash, err := bcrypt.GenerateFromPassword([]byte(body.MotDePasse), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Error hashing password", http.StatusInternalServerError)
			return
		}

		err = service.AddUser(body.Nom, body.Email, string(hash), body.Role)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
			response := struct {
				Meta config.JsonMeta `json:"meta"`
			}{
				Meta: config.JsonMeta{
					Status:  200,
					Message: "success",
				},
			}

			config.EnableCORS(w)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		}
}