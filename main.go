package main

import (
	// import du package prices qui encapsule la logique de calcul TTC
	"example.com/price-calculator/prices"
)

func main() {
	// La logique de calcul a été extraite dans le package prices.
	// main.go se contente désormais d'orchestrer les traitements :
	// pour chaque taux de taxe, on crée un job et on l'exécute.
	taxRates := []float64{0, 0.021, 0.055, 0.1, 0.2}

	for _, taxRate := range taxRates {
		// crée un job pré-configuré avec les prix HT par défaut et le taux courant
		priceJob := prices.NewTaxIncludedPriceJob(taxRate)
		// calcule les prix TTC et affiche le résultat pour ce taux
		priceJob.Process()
	}
}
