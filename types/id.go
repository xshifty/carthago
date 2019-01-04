package types

import (
	uuid "github.com/satori/go.uuid"
)

type ID string

func NewID() ID {
	return ID(uuid.NewV4().String())
}

func (innerID ID) SameAs(outterID ID) bool {
	return innerID == outterID
}
