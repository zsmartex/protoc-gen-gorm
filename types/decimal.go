package types

import "github.com/shopspring/decimal"

func FromDecimal(d decimal.Decimal) *Decimal {
	return &Decimal{
		Value: d.CoefficientInt64(),
		Scale: d.Exponent(),
	}
}

func (d *Decimal) ToDecimal() decimal.Decimal {
	return decimal.New(d.Value, d.Scale)
}
