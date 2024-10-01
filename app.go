package main

import ("fmt" "example.com/practise/prices")

func main() {
	prices := []float64{10, 20, 30}
	taxRates := []float64{0.0, 0.07, 0.1, 0.2}

	mappings := make(map[float64][]float64)

	for _, taxRate := range taxRates {
		priceJob:=prices.NewTaxIncludedPriceJob(taxRate)
		priceJob.Process()
	}

	fmt.Println(mappings)

}
