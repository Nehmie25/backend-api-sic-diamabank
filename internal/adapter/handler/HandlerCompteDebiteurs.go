package handler

import(
	"encoding/xml"
	"net/http"
	"GoWebapitest/config"
	"GoWebapitest/internal/core/service"
)

func HandlerCompteDebiteurs(service *service.CompteDebiteursService) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {

		date := r.URL.Query().Get("date")
		if date == "" {
			http.Error(w, "date requise", http.StatusBadRequest)
			return
		}

		declaration, err := service.BuildDeclarationEncours(date)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := config.Response{
			Meta: config.Meta{
				Status: 200,
				Message: "success",
			},
			Data: config.Data{Declaration: declaration},
		}

		config.EnableCORS(w)
		w.Header().Set("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(response)
	}
}