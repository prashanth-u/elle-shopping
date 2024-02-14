package routes

import (
    "net/http"
	"encoding/json"
	"crypto/rand"
	"encoding/base64"
	"shopping/internal/service"
	"shopping/internal/thirdparty"
)

func generateToken() (string, error) {
    b := make([]byte, 32)
    _, err := rand.Read(b)
    if err != nil {
        return "", err
    }
    return base64.URLEncoding.EncodeToString(b), nil
}

func LoginRoutes(mux *http.ServeMux, userService *service.UserService) {
    mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}
		response := map[string]string{"message": "hello login"}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonResponse)
    })
	mux.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		firstName := r.Form.Get("firstName")
		lastName := r.Form.Get("lastName")
		token, err := generateToken()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		ctx := r.Context()
		userService.Register(ctx, username, password, firstName, lastName, token)
		thirdparty.SendVerificationEmail(username, token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response := map[string]bool{"status": true}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonResponse)
	})
}