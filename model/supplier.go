package model

import (
	types "github.com/xshifty/carthago/types"
)

type Supplier struct {
	id   types.ID
	name string
}

func NewSupplier(id types.ID, name string) *Supplier {
	return &Supplier{
		id:   id,
		name: name,
	}
}

func (supplier *Supplier) ID() types.ID {
	return supplier.id
}

func (supplier *Supplier) Name() string {
	return supplier.name
}
