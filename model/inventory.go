package model

import (
	types "github.com/xshifty/carthago/types"
)

type Inventory struct {
	product     *Product
	quantity    uint32
	averageCost *types.Money
}

func NewInventory(product *Product, quantity uint32, averageCost *types.Money) *Inventory {
	return &Inventory{
		product:     product,
		quantity:    quantity,
		averageCost: averageCost,
	}
}
