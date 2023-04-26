package pricing

import (
	"github.com/shopspring/decimal"
)

type Price struct {
	Value        decimal.Decimal
	CurrencyCode string
}

func New(value string, currencyCode string) (p Price, err error) {
	price, err := decimal.NewFromString(value)
	if err != nil {
		return
	}

	// TODO: add taxes processing

	p = Price{
		Value:        price,
		CurrencyCode: currencyCode,
	}
	return
}
