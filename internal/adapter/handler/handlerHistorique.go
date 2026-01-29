package handler

import (
	"GoWebapitest/config"
	"GoWebapitest/internal/core/domain/models"
	"GoWebapitest/internal/core/service"
	"encoding/json"
	"net/http"
	"strconv"
)


func HandlerHistorique(service *service.HistoriqueService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		page := r.URL.Query().Get("page")


		pageInt, err := strconv.Atoi(page)
		if err != nil {
			pageInt = 1
		}

		if pageInt < 1 {
			pageInt = 1
		}

		historiques, pagination, err := service.GetAllLogs(pageInt)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération des historiques", http.StatusInternalServerError)
			return
		}

		response := struct {
			Meta config.JsonMeta `json:"meta"`
			Data []*models.Historique
			Pagination *models.Pagination
		}{
			Meta: config.JsonMeta{
				Status:  200,
				Message: "success",
			},
			Data: historiques,
			Pagination: pagination,
		}

		config.EnableCORS(w)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}