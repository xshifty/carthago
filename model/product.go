package model

import (
	types "github.com/xshifty/carthago/types"
)

type Product struct {
	id          types.ID
	description string
}

func NewProduct(id types.ID, description string) *Product {
	return &Product{
		id:          id,
		description: description,
	}
}

func (product *Product) ID() types.ID {
	return product.id
}

func (product *Product) Description() string {
	return product.description
}

type ProductList []*Product
