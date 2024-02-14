package routes

import (
    "net/http"
	"encoding/json"
)

func UserRoutes(mux *http.ServeMux) {
    mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]string{"message": "hello user"}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonResponse)
    })
}