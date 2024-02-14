package routes

import (
    "net/http"
	"encoding/json"
	"shopping/internal/service"
	"shopping/internal/models"
)

func ProductRoutes(mux *http.ServeMux, productService *service.ProductService) {
    mux.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		response, err := productService.GetProducts(ctx)
		jsonResponse, err := json.Marshal(response)
		requestID := r.Context().Value("requestID").(string)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonResponse)
    })
	mux.HandleFunc("/products/add", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var product models.Product

		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		response, err := productService.AddProduct(ctx, product)
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonResponse)
    })
}