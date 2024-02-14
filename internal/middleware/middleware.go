package middleware

import (
    "net/http"
    "shopping/internal/service"
    "shopping/internal/routes"
    "log"
    "context"
    "shopping/internal/models"
    "github.com/google/uuid"
)

func ValidRouteMiddleware(next http.Handler, userService *service.UserService) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        
        if _, ok := routes.ValidRoutes[r.URL.Path]; !ok {
            http.NotFound(w, r)
            return
        }
        if r.URL.Path == "/login" || r.URL.Path == "/register" {
            log.Println("Login/Register")
            next.ServeHTTP(w, r)
            return
        } else {
            ctx := context.Background()
            params := r.URL.Query()
            sessionId := params.Get("sessionId")
            if sessionId == "" {
                http.Error(w, "Unauthorized", http.StatusUnauthorized)
                return
            }
            log.Println("SessionId: ", sessionId)
            user, err := userService.ValidateRequest(ctx, sessionId)     
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }          
            if user == (models.User{}) {
                http.Error(w, "Unauthorized", http.StatusUnauthorized)
                return
            }
        }
        next.ServeHTTP(w, r)
    })
}

func TraceRequestMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        requestID := uuid.New().String()
        ctx := context.WithValue(r.Context(), "requestID", requestID)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

func AddServerTimeOutMiddleWare(next http.Handler) http.Handler {
    return http.TimeoutHandler(next, 5, "Server Timeout")
}