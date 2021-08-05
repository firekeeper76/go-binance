package common

import (
	"github.com/shopspring/decimal"
)

// PriceLevel is a common structure for bids and asks in the
// order book.
type PriceLevel struct {
	Price    string
	Quantity string
}

// Parse parses this PriceLevel's Price and Quantity and
// returns them both.  It also returns an error if either
// fails to parse.
func (p *PriceLevel) Parse() (decimal.Decimal, decimal.Decimal, error) {
	price, err := decimal.NewFromString(p.Price)
	if err != nil {
		return decimal.Zero, decimal.Zero, err
	}
	quantity, err := decimal.NewFromString(p.Quantity)
	if err != nil {
		return decimal.Zero, decimal.Zero, err
	}
	return price, quantity, nil
}
