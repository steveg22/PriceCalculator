package main

import (
	"fmt"

	"example.com/practise/cmdmanager"
	//	"example.com/practise/filemanager"
	"example.com/practise/prices"
)

func main() {
	taxRates := []float64{0.0, 0.07, 0.1, 0.2}

	for _, taxRate := range taxRates {
		//fm := filemanager.New("prices.txt", fmt.Sprintf("result%.0f.json", taxRate*100))
		cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(cmdm, taxRate)
		err := priceJob.Process()
		if err != nil {
			fmt.Println("Could not process job")
			fmt.Println(err)
		}
	}
}
