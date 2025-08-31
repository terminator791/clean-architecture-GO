package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/terminator791/clean-architecture-GO/internal/domain/entity"
)

// ProductRepository defines the interface for product data operations
type ProductRepository interface {
	// Create creates a new product
	Create(ctx context.Context, product *entity.Product) error
	
	// GetByID retrieves a product by its ID
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Product, error)
	
	// GetAll retrieves all products with optional pagination
	GetAll(ctx context.Context, limit, offset int) ([]*entity.Product, error)
	
	// Update updates an existing product
	Update(ctx context.Context, product *entity.Product) error
	
	// Delete deletes a product by its ID
	Delete(ctx context.Context, id uuid.UUID) error
}