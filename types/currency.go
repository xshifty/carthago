package types

type Currency uint16

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

type Money struct {
	value    float32
	currency Currency
}

func NewMoney(value float32, currency Currency) *Money {
	return &Money{
		value:    value,
		currency: currency,
	}
}

func (money *Money) Value() float32 {
	return money.value
}

func (money *Money) Currency() Currency {
	return money.currency
}
