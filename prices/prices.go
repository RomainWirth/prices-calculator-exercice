// Package prices fournit les outils pour calculer des prix TTC
// à partir d'un taux de taxe et d'une liste de prix HT.
package prices

import (
	// lecture ligne par ligne du fichier
	"fmt"
	// accès au système de fichiers
	// conversion des lignes texte en float64
	"example.com/price-calculator/conversion"
	"example.com/price-calculator/filemanager"
)

// TaxIncludedPriceJob représente un travail de calcul de prix TTC.
// Il regroupe le taux de taxe, les prix HT en entrée et la map des prix TTC calculés.
type TaxIncludedPriceJob struct {
	TaxRate           float64           // taux de taxe à appliquer (ex: 0.2 pour 20%)
	InputPrices       []float64         // liste des prix HT à traiter
	TaxIncludedPrices map[string]string // résultats : clé = prix HT formaté, valeur = prix TTC
}

func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := filemanager.ReadLines("prices.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		fmt.Println(err)
		return
	}

	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()

	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result

	filemanager.WriteJSON(fmt.Sprintf("./results_json/tax_included_prices_%.2f.json", job.TaxRate*100), job)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30}, // prix HT par défaut
		TaxRate:     taxRate,
	}
}
