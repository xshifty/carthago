package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type ID string

// NewID create a new ID string
func NewID() ID {
	return ID(uuid.NewV4().String())
}

// SameAs compare two ID's
func (innerID ID) SameAs(outterID ID) bool {
	return innerID == outterID
}

// UnixTimestamp type definition
type UnixTimestamp int64

// NewUnixTimestamp return the current time as a Timestamp
func NewUnixTimestamp() UnixTimestamp {
	return UnixTimestamp(time.Now().Unix())
}

// DatetimeString UnixTimestamp function which return a string in DateTime format
func (timestamp UnixTimestamp) DatetimeString() string {
	return time.Unix(int64(timestamp), 0).String()
}

// Currency type representation
type Currency uint16

// Currency possible values
const (
	BRL Currency = iota
	USD
	EUR
)

func (currency Currency) String() string {
	switch currency {
	case BRL:
		return "BRL"
	case USD:
		return "USD"
	case EUR:
		return "EUR"
	}

	return ""
}

// Money representation
type Money struct {
	value    float32
	currency Currency
}

// NewMoney build a new Money struct
func NewMoney(value float32, currency Currency) *Money {
	return &Money{
		value:    value,
		currency: currency,
	}
}

// Value from a Money
func (money *Money) Value() float32 {
	return money.value
}

// Currency from a Money
func (money *Money) Currency() Currency {
	return money.currency
}

// DimensionUnit represent a measure unit for a product dimension
type DimensionUnit uint16

// Possible DimensionUnit values
const (
	METER DimensionUnit = iota
	CENTIMETER
	MILIMITER
)

// Dimension definition. Represents the dimensions of a product
type Dimension struct {
	width  float32
	height float32
	length float32
	unit   DimensionUnit
}

// NewDimension builds a new Dimension struct
func NewDimension(width float32, height float32, length float32, unit DimensionUnit) *Dimension {
	return &Dimension{
		width:  width,
		height: height,
		length: length,
		unit:   unit,
	}
}

// Width from a Dimension
func (dimension *Dimension) Width() float32 {
	return dimension.width
}

// Height from a Dimension
func (dimension *Dimension) Height() float32 {
	return dimension.height
}

// Length from a Dimension
func (dimension *Dimension) Length() float32 {
	return dimension.length
}

// Unit type from a Dimension
func (dimension *Dimension) Unit() DimensionUnit {
	return dimension.unit
}
