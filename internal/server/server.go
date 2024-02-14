package server

import (
    "net/http"
    "shopping/internal/routes"
    "shopping/internal/middleware"
    "shopping/internal/repository"
    "shopping/internal/service"
    "database/sql"
)

type Server struct {
    httpServer *http.Server
}

func NewServer(port string, db *sql.DB) *Server {
    productRepo := repository.NewProductRepository(db)
    productService := service.NewProductService(productRepo)
    userRepo := repository.NewUserRepository(db)
    userService := service.NewUserService(userRepo)

	mux := http.NewServeMux()
    routes.UserRoutes(mux)
    routes.ProductRoutes(mux, productService)
    routes.CategoryRoutes(mux)
    routes.LoginRoutes(mux, userService)
    validatedMux := middleware.ValidRouteMiddleware(mux, userService)
    tracedMux := middleware.TraceRequestMiddleware(validatedMux)
    s := &Server{
        httpServer: &http.Server{
            Addr: port,
			Handler: tracedMux,
        },
    }
    return s
}

func (s *Server) Start() error {
    return s.httpServer.ListenAndServe()
}