package usecase

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/terminator791/clean-architecture-GO/internal/domain/entity"
	"github.com/terminator791/clean-architecture-GO/internal/domain/repository"
)

var (
	ErrProductNotFound = errors.New("product not found")
	ErrInvalidInput    = errors.New("invalid input")
)

// ProductUseCase handles product business logic
type ProductUseCase struct {
	repo repository.ProductRepository
}

// NewProductUseCase creates a new product use case
func NewProductUseCase(repo repository.ProductRepository) *ProductUseCase {
	return &ProductUseCase{
		repo: repo,
	}
}

// CreateProduct creates a new product
func (uc *ProductUseCase) CreateProduct(ctx context.Context, name, description string, price float64) (*entity.Product, error) {
	if name == "" {
		return nil, ErrInvalidInput
	}
	if price < 0 {
		return nil, ErrInvalidInput
	}

	product := entity.NewProduct(name, description, price)
	
	if err := uc.repo.Create(ctx, product); err != nil {
		return nil, err
	}

	return product, nil
}

// GetProduct retrieves a product by ID
func (uc *ProductUseCase) GetProduct(ctx context.Context, id uuid.UUID) (*entity.Product, error) {
	product, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, ErrProductNotFound
	}
	return product, nil
}

// GetAllProducts retrieves all products with pagination
func (uc *ProductUseCase) GetAllProducts(ctx context.Context, limit, offset int) ([]*entity.Product, error) {
	if limit <= 0 {
		limit = 10 // default limit
	}
	if offset < 0 {
		offset = 0
	}

	return uc.repo.GetAll(ctx, limit, offset)
}

// UpdateProduct updates an existing product
func (uc *ProductUseCase) UpdateProduct(ctx context.Context, id uuid.UUID, name, description string, price float64) (*entity.Product, error) {
	if name == "" {
		return nil, ErrInvalidInput
	}
	if price < 0 {
		return nil, ErrInvalidInput
	}

	product, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, ErrProductNotFound
	}

	product.Update(name, description, price)

	if err := uc.repo.Update(ctx, product); err != nil {
		return nil, err
	}

	return product, nil
}

// DeleteProduct deletes a product by ID
func (uc *ProductUseCase) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	product, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if product == nil {
		return ErrProductNotFound
	}

	return uc.repo.Delete(ctx, id)
}