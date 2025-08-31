package router

import (
	"net/http"
	"strings"

	"github.com/terminator791/clean-architecture-GO/internal/adapter/controller"
)

// Router handles HTTP routing
type Router struct {
	productController *controller.ProductController
}

// NewRouter creates a new router
func NewRouter(productController *controller.ProductController) *Router {
	return &Router{
		productController: productController,
	}
}

// ServeHTTP implements http.Handler
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Enable CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// Handle preflight requests
	if req.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	path := req.URL.Path

	// Health check endpoint
	if path == "/health" && req.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"healthy"}`))
		return
	}

	// Product routes
	if strings.HasPrefix(path, "/products") {
		r.handleProductRoutes(w, req)
		return
	}

	// Not found
	http.NotFound(w, req)
}

// handleProductRoutes handles product-related routes
func (r *Router) handleProductRoutes(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path

	switch {
	case path == "/products" && req.Method == http.MethodPost:
		r.productController.CreateProduct(w, req)
	case path == "/products" && req.Method == http.MethodGet:
		r.productController.GetAllProducts(w, req)
	case strings.HasPrefix(path, "/products/") && req.Method == http.MethodGet:
		r.productController.GetProduct(w, req)
	case strings.HasPrefix(path, "/products/") && req.Method == http.MethodPut:
		r.productController.UpdateProduct(w, req)
	case strings.HasPrefix(path, "/products/") && req.Method == http.MethodDelete:
		r.productController.DeleteProduct(w, req)
	default:
		http.NotFound(w, req)
	}
}