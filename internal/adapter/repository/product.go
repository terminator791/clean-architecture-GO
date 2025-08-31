package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/terminator791/clean-architecture-GO/internal/domain/entity"
)

// PostgresProductRepository implements ProductRepository interface
type PostgresProductRepository struct {
	db *sqlx.DB
}

// NewPostgresProductRepository creates a new PostgreSQL product repository
func NewPostgresProductRepository(db *sqlx.DB) *PostgresProductRepository {
	return &PostgresProductRepository{
		db: db,
	}
}

// Create creates a new product in the database
func (r *PostgresProductRepository) Create(ctx context.Context, product *entity.Product) error {
	query := `
		INSERT INTO products (id, name, description, price, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := r.db.ExecContext(ctx, query,
		product.ID,
		product.Name,
		product.Description,
		product.Price,
		product.CreatedAt,
		product.UpdatedAt,
	)
	return err
}

// GetByID retrieves a product by its ID
func (r *PostgresProductRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.Product, error) {
	var product entity.Product
	query := `
		SELECT id, name, description, price, created_at, updated_at
		FROM products
		WHERE id = $1`

	err := r.db.GetContext(ctx, &product, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &product, nil
}

// GetAll retrieves all products with pagination
func (r *PostgresProductRepository) GetAll(ctx context.Context, limit, offset int) ([]*entity.Product, error) {
	var products []*entity.Product
	query := `
		SELECT id, name, description, price, created_at, updated_at
		FROM products
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2`

	err := r.db.SelectContext(ctx, &products, query, limit, offset)
	if err != nil {
		return nil, err
	}
	return products, nil
}

// Update updates an existing product
func (r *PostgresProductRepository) Update(ctx context.Context, product *entity.Product) error {
	query := `
		UPDATE products
		SET name = $2, description = $3, price = $4, updated_at = $5
		WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query,
		product.ID,
		product.Name,
		product.Description,
		product.Price,
		product.UpdatedAt,
	)
	return err
}

// Delete deletes a product by its ID
func (r *PostgresProductRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}