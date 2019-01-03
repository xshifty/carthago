package model

type Supplier struct {
	id   ID
	name string
}

func NewSupplier(id ID, name string) *Supplier {
	return &Supplier{
		id:   id,
		name: name,
	}
}

func (supplier *Supplier) ID() ID {
	return supplier.id
}

func (supplier *Supplier) Name() string {
	return supplier.name
}
