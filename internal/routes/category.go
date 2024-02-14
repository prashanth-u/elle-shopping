package routes

import (
    "net/http"
	"encoding/json"
)

func CategoryRoutes(mux *http.ServeMux) {
    mux.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]string{"message": "hello categories"}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonResponse)
    })
}