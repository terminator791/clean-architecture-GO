package entity

import (
	"time"

	"github.com/google/uuid"
)

// Product represents a product entity
type Product struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Price       float64   `json:"price" db:"price"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// NewProduct creates a new product with generated UUID
func NewProduct(name, description string, price float64) *Product {
	now := time.Now()
	return &Product{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
		Price:       price,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

// Update updates the product fields and UpdatedAt timestamp
func (p *Product) Update(name, description string, price float64) {
	p.Name = name
	p.Description = description
	p.Price = price
	p.UpdatedAt = time.Now()
}