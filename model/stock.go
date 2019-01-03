package model

type AcquiredBatch struct {
	id        ID
	supplier  *Supplier
	product   *Product
	cost      *Money
	quantity  uint32
	createdAt UnixTimestamp
}

func NewAcquiredBatch(id ID, supplier *Supplier, product *Product, cost *Money, quantity uint32, createdAt UnixTimestamp) *AcquiredBatch {
	return &AcquiredBatch{
		id:        id,
		supplier:  supplier,
		product:   product,
		cost:      cost,
		quantity:  quantity,
		createdAt: createdAt,
	}
}

func (batch *AcquiredBatch) ID() ID {
	return batch.id
}

func (batch *AcquiredBatch) Supplier() *Supplier {
	return batch.supplier
}

func (batch *AcquiredBatch) Product() *Product {
	return batch.product
}

func (batch *AcquiredBatch) Cost() *Money {
	return batch.cost
}

func (batch *AcquiredBatch) Quantity() uint32 {
	return batch.quantity
}

func (batch *AcquiredBatch) CreatedAt() UnixTimestamp {
	return batch.createdAt
}

type SellingOrderStatus uint16

const (
	SELLING_ORDER_STARTED SellingOrderStatus = iota
	SELLING_ORDER_CANCELED
	SELLING_ORDER_CONCLUDED
)

type SellingOrder struct {
	id        ID
	product   *Product
	price     *Money
	quantity  uint32
	status    SellingOrderStatus
	createdAt UnixTimestamp
}

func NewSellingOrder(id ID, product *Product, price *Money, quantity uint32, status SellingOrderStatus, createdAt UnixTimestamp) *SellingOrder {
	return &SellingOrder{
		id:        id,
		product:   product,
		price:     price,
		quantity:  quantity,
		status:    status,
		createdAt: createdAt,
	}
}
