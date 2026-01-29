package handler

import (
	"encoding/json"
	"net/http"

	"GoWebapitest/config"
	"GoWebapitest/internal/core/service"

	"golang.org/x/crypto/bcrypt"
)

func HandlerResetPassword(service *service.ResetPasswordService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//recupérer userId depuis le body
		var body struct {
			UserId   string `json:"userId"`
			Password string `json:"motdepasse"`
		}
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Error hashing password", http.StatusInternalServerError)
			return
		}

		err = service.ResetPassword(body.UserId, string(hash))
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
