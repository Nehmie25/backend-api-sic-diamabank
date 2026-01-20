package handler

import (
	"GoWebapitest/config"
	"GoWebapitest/internal/core/domain/models"
	"GoWebapitest/internal/core/service"
	"encoding/json"
	"fmt"
	"net/http"
)

func ModifyStatus(service *service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Récupérer les paramètres du formulaire par POST

		var body models.UserStatus

		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			fmt.Println(&body)
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		fmt.Println(&body)
		err = service.ModifyStatus(body.Id, body.Isactive)
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
