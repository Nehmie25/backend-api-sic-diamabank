package handler

import (
	"GoWebapitest/config"
	"GoWebapitest/internal/core/service"
	"encoding/json"
	"net/http"
	"strconv"
)

func ModifyStatus(service *service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Récupérer les paramètres du formulaire par POST
		idStr := r.FormValue("id")
		isactiveStr := r.FormValue("isactive")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid id parameter", http.StatusBadRequest)
			return
		}

		isactive, err := strconv.ParseBool(isactiveStr)
		if err != nil {
			http.Error(w, "Invalid isactive parameter", http.StatusBadRequest)
			return
		}

		err = service.ModifyStatus(id, isactive)

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