package repository

import (
	"context"
	"errors"
)

var (
	// ErrNotFound is returned when a resource is not found
	ErrNotFound = errors.New("resource not found")
	// ErrAlreadyExists is returned when a resource already exists
	ErrAlreadyExists = errors.New("resource already exists")
)

// Repository represents data access layer
type Repository struct {
	// Add your database connection or other storage dependencies here
}

// NewRepository creates a new repository instance
func NewRepository() *Repository {
	return &Repository{}
}

// Example method - replace with your actual data access logic
func (r *Repository) GetByID(ctx context.Context, id string) (interface{}, error) {
	// Add your data access logic here
	return nil, ErrNotFound
}
