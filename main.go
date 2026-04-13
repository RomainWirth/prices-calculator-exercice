package main

import (
	"example.com/price-calculator/cmdmanager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.021, 0.055, 0.1, 0.2}

	for _, taxRate := range taxRates {
		// fm := filemanager.New("./data/prices.txt", fmt.Sprintf("./results_json/tax_included_prices_%.2f.json", taxRate*100))
		cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(cmdm, taxRate)
		priceJob.Process()
	}
}
