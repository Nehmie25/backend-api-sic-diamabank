package handler

import (
	"encoding/json"
	"net/http"
	"GoWebapitest/config"
	"GoWebapitest/internal/core/service"
)

func HandlerAddUser(service *service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Récupérer les paramètres du formulaire par POST
		nom := r.FormValue("nom")
		email := r.FormValue("email")
		motdepasse := r.FormValue("motdepasse")
		role := r.FormValue("role")

		err := service.AddUser(nom, email, motdepasse, role)
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