package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/terminator791/clean-architecture-GO/internal/usecase"
)

// ProductController handles HTTP requests for products
type ProductController struct {
	useCase *usecase.ProductUseCase
}

// NewProductController creates a new product controller
func NewProductController(useCase *usecase.ProductUseCase) *ProductController {
	return &ProductController{
		useCase: useCase,
	}
}

// CreateProductRequest represents the request body for creating a product
type CreateProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

// UpdateProductRequest represents the request body for updating a product
type UpdateProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// writeErrorResponse writes an error response to the HTTP response writer
func writeErrorResponse(w http.ResponseWriter, statusCode int, err string, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{
		Error:   err,
		Message: message,
	})
}

// writeJSONResponse writes a JSON response to the HTTP response writer
func writeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// CreateProduct handles POST /products
func (c *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "INVALID_JSON", "Invalid JSON format")
		return
	}

	product, err := c.useCase.CreateProduct(r.Context(), req.Name, req.Description, req.Price)
	if err != nil {
		if err == usecase.ErrInvalidInput {
			writeErrorResponse(w, http.StatusBadRequest, "INVALID_INPUT", err.Error())
			return
		}
		writeErrorResponse(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Failed to create product")
		return
	}

	writeJSONResponse(w, http.StatusCreated, product)
}

// GetProduct handles GET /products/{id}
func (c *ProductController) GetProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/products/"):]
	id, err := uuid.Parse(idStr)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "INVALID_UUID", "Invalid product ID format")
		return
	}

	product, err := c.useCase.GetProduct(r.Context(), id)
	if err != nil {
		if err == usecase.ErrProductNotFound {
			writeErrorResponse(w, http.StatusNotFound, "PRODUCT_NOT_FOUND", "Product not found")
			return
		}
		writeErrorResponse(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Failed to get product")
		return
	}

	writeJSONResponse(w, http.StatusOK, product)
}

// GetAllProducts handles GET /products
func (c *ProductController) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	limit := 10 // default
	offset := 0 // default

	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	if offsetStr != "" {
		if o, err := strconv.Atoi(offsetStr); err == nil && o >= 0 {
			offset = o
		}
	}

	products, err := c.useCase.GetAllProducts(r.Context(), limit, offset)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Failed to get products")
		return
	}

	writeJSONResponse(w, http.StatusOK, products)
}

// UpdateProduct handles PUT /products/{id}
func (c *ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/products/"):]
	id, err := uuid.Parse(idStr)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "INVALID_UUID", "Invalid product ID format")
		return
	}

	var req UpdateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "INVALID_JSON", "Invalid JSON format")
		return
	}

	product, err := c.useCase.UpdateProduct(r.Context(), id, req.Name, req.Description, req.Price)
	if err != nil {
		if err == usecase.ErrProductNotFound {
			writeErrorResponse(w, http.StatusNotFound, "PRODUCT_NOT_FOUND", "Product not found")
			return
		}
		if err == usecase.ErrInvalidInput {
			writeErrorResponse(w, http.StatusBadRequest, "INVALID_INPUT", err.Error())
			return
		}
		writeErrorResponse(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Failed to update product")
		return
	}

	writeJSONResponse(w, http.StatusOK, product)
}

// DeleteProduct handles DELETE /products/{id}
func (c *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/products/"):]
	id, err := uuid.Parse(idStr)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "INVALID_UUID", "Invalid product ID format")
		return
	}

	err = c.useCase.DeleteProduct(r.Context(), id)
	if err != nil {
		if err == usecase.ErrProductNotFound {
			writeErrorResponse(w, http.StatusNotFound, "PRODUCT_NOT_FOUND", "Product not found")
			return
		}
		writeErrorResponse(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Failed to delete product")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}