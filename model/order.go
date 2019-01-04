package model

import (
	types "github.com/xshifty/carthago/types"
)

const (
	ORDER_STARTED uint16 = iota
	ORDER_CANCELED
	ORDER_CONCLUDED
)

type Order struct {
	id        types.ID
	product   *Product
	price     *types.Money
	quantity  uint32
	status    uint16
	createdAt types.UnixTimestamp
}

func NewOrder(id types.ID, product *Product, price *types.Money, quantity uint32, status uint16, createdAt types.UnixTimestamp) *Order {
	return &Order{
		id:        id,
		product:   product,
		price:     price,
		quantity:  quantity,
		status:    status,
		createdAt: createdAt,
	}
}

func (order *Order) ID() types.ID {
	return order.id
}

func (order *Order) Product() *Product {
	return order.product
}

func (order *Order) Price() *types.Money {
	return order.price
}
