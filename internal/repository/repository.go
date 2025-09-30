package repository

import (
	"context"
	"errors"
)

var (
	ErrNotFound      = errors.New("resource not found")
	ErrAlreadyExists = errors.New("resource already exists")
)

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) GetByID(_ context.Context, _ string) (interface{}, error) {
	return nil, ErrNotFound
}
