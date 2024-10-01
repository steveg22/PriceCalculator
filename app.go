package main

import (
	"example.com/practise/prices"
)

func main() {
	taxRates := []float64{0.0, 0.07, 0.1, 0.2}

	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJob(taxRate)
		priceJob.Process()
	}
}
