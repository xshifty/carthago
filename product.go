package carthago

const (
    UNIT Measure = iota
)

type (
    Product struct {
        ID string
        Description string
        Unit Measure
    }

    Measure int
)

func NewProduct(id string, description string, unit Measure) (Product, error) {
    return Product{ID: id, Description: description, Unit: unit}, nil
}

func (measure Measure) String() string {
    switch measure {
    case UNIT:
        return "unit"
    }

    return ""
}
