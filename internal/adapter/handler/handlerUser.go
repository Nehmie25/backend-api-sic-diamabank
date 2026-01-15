package handler

import (
	"encoding/json"
	"net/http"

	"GoWebapitest/config"
	"GoWebapitest/internal/core/service"
)

func HandlerUser(service *service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		Users, err := service.GetAllUsers()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response := struct {
			Meta config.JsonMeta `json:"meta"`
			Data config.JsonData `json:"data"`
		}{
			Meta: config.JsonMeta{
				Status:  200,
				Message: "success",
			},
			Data: config.JsonData{Users: Users},
		}

		config.EnableCORS(w)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
