package prices

import (
	"fmt"

	"example.com/practise/conversion"
	"example.com/practise/iomanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
	IOManager         iomanager.IOManager `json:"-"`
}

func (job *TaxIncludedPriceJob) LoadData() error {
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		return err
	}

	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		return err
	}

	job.InputPrices = prices
	return nil
}

func (job *TaxIncludedPriceJob) Process() error {
	err := job.LoadData()
	if err != nil {
		return err
	}

	mappings := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := fmt.Sprintf("%.2f", price*(1+job.TaxRate))
		mappings[fmt.Sprintf("%.2f", price)] = taxIncludedPrice
	}

	job.TaxIncludedPrices = mappings

	fmt.Println(job.TaxRate, job.TaxIncludedPrices)

	return job.IOManager.WriteResult(job)
}

func NewTaxIncludedPriceJob(fm iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
		IOManager:   fm,
	}
}
