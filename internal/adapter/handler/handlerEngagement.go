package handler

import (
	"GoWebapitest/config"
	"GoWebapitest/internal/core/service"
	"encoding/xml"
	"net/http"
)

func HandlerEngagement(service *service.EngagementsService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		date := r.URL.Query().Get("date")
		if date == "" {
			http.Error(w, "date requise", http.StatusBadRequest)
			return
		}

		declaration, err := service.BuildDeclarationEngagements(date)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := config.Response{
			Meta: config.Meta{
				Status:  200,
				Message: "success",
			},
			Data: config.Data{Declaration: declaration},
		}

		config.EnableCORS(w)

		xml.NewEncoder(w).Encode(response)
	}
}
