package model

import (
	types "github.com/xshifty/carthago/types"
)

type Batch struct {
	id        types.ID
	supplier  *Supplier
	product   *Product
	cost      *types.Money
	quantity  uint32
	createdAt types.UnixTimestamp
}

func NewBatch(id types.ID, supplier *Supplier, product *Product, cost *types.Money, quantity uint32, createdAt types.UnixTimestamp) *Batch {
	return &Batch{
		id:        id,
		supplier:  supplier,
		product:   product,
		cost:      cost,
		quantity:  quantity,
		createdAt: createdAt,
	}
}

func (batch *Batch) ID() types.ID {
	return batch.id
}

func (batch *Batch) Supplier() *Supplier {
	return batch.supplier
}

func (batch *Batch) Product() *Product {
	return batch.product
}

func (batch *Batch) Cost() *types.Money {
	return batch.cost
}

func (batch *Batch) Quantity() uint32 {
	return batch.quantity
}

func (batch *Batch) CreatedAt() types.UnixTimestamp {
	return batch.createdAt
}

const (
	SELLING_ORDER_STARTED uint16 = iota
	SELLING_ORDER_CANCELED
	SELLING_ORDER_CONCLUDED
)

type SellingOrder struct {
	id        types.ID
	product   *Product
	price     *types.Money
	quantity  uint32
	status    uint16
	createdAt types.UnixTimestamp
}

func NewSellingOrder(id types.ID, product *Product, price *types.Money, quantity uint32, status uint16, createdAt types.UnixTimestamp) *SellingOrder {
	return &SellingOrder{
		id:        id,
		product:   product,
		price:     price,
		quantity:  quantity,
		status:    status,
		createdAt: createdAt,
	}
}
