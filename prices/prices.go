// Package prices fournit les outils pour calculer des prix TTC
// à partir d'un taux de taxe et d'une liste de prix HT.
package prices

import (
	"fmt"

	"example.com/price-calculator/conversion"
	"example.com/price-calculator/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"`                   // gestionnaire de fichiers pour lire les prix HT et écrire les résultats
	TaxRate           float64             `json:"tax_rate"`            // taux de taxe à appliquer (ex: 0.2 pour 20%)
	InputPrices       []float64           `json:"input_prices"`        // liste des prix HT à traiter
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"` // résultats : clé = prix HT formaté, valeur = prix TTC
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

	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result

	return job.IOManager.WriteResult(job)
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   iom,
		InputPrices: []float64{10, 20, 30}, // prix HT par défaut
		TaxRate:     taxRate,
	}
}
