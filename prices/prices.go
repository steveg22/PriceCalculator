package prices

import (
	"example.com/practise/conversion"
	"example.com/practise/filemanager"
	"fmt"
)

const priceFile = "prices.txt"

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]string
}

func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := filemanager.ReadLines(priceFile)
	if err != nil {
		fmt.Println("Error Reading file")
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		fmt.Println("Error converting to number")
		fmt.Println(err)
		return
	}

	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	mappings := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := fmt.Sprintf("%.2f", price*(1+job.TaxRate))
		mappings[fmt.Sprintf("%.2f", price)] = taxIncludedPrice
	}

	job.TaxIncludedPrices = mappings

	fmt.Println(job.TaxRate, job.TaxIncludedPrices)

	filemanager.WriteJSON(fmt.Sprintf("result_%.0f.json", job.TaxRate*100), job)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
