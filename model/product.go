package model

import (
	"errors"
)

var (
	ErrProductNotFound = errors.New("Product not found.")
)

type Product struct {
	id          ID
	description string
}

func NewProduct(id ID, description string) *Product {
	return &Product{
		id:          id,
		description: description,
	}
}

func (product *Product) ID() ID {
	return product.id
}

func (product *Product) Description() string {
	return product.description
}

type ProductList []*Product
