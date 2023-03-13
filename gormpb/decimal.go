package gormpb

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

func (d *Decimal) ToNullDecimal() decimal.NullDecimal {
	return decimal.NewNullDecimal(decimal.New(d.Value, d.Scale))
}

func FromNullDecimal(d decimal.NullDecimal) *Decimal {
	return &Decimal{
		Value: d.Decimal.CoefficientInt64(),
		Scale: d.Decimal.Exponent(),
	}
}
