package operations

import "github.com/shopspring/decimal"

func SafeSum(n ...float64) decimal.Decimal {
	var res decimal.Decimal
	for _, f := range n {
		res = res.Add(decimal.NewFromFloat(f))
	}
	return res
}
