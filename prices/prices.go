// Package prices fournit les outils pour calculer des prix TTC
// à partir d'un taux de taxe et d'une liste de prix HT.
package prices

import (
	"bufio" // lecture ligne par ligne du fichier
	"fmt"
	"os" // accès au système de fichiers
	"strconv"
)

// TaxIncludedPriceJob représente un travail de calcul de prix TTC.
// Il regroupe le taux de taxe, les prix HT en entrée et la map des prix TTC calculés.
type TaxIncludedPriceJob struct {
	TaxRate           float64            // taux de taxe à appliquer (ex: 0.2 pour 20%)
	InputPrices       []float64          // liste des prix HT à traiter
	TaxIncludedPrices map[string]float64 // résultats : clé = prix HT formaté, valeur = prix TTC
}

// LoadData lit le fichier "prices.txt" ligne par ligne et charge les prix HT dans le job.
// Chaque ligne du fichier doit contenir un prix (format texte).
// En cas d'erreur d'ouverture ou de lecture, un message est affiché et la fonction retourne sans modifier le job.
func (job *TaxIncludedPriceJob) LoadData() {
	// ouverture du fichier contenant les prix HT
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return
	}

	// scanner pour lire le fichier ligne par ligne
	scanner := bufio.NewScanner(file)

	var lines []string

	// lecture de chaque ligne jusqu'à la fin du fichier
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// vérification d'une éventuelle erreur survenue pendant le scan
	err = scanner.Err()
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
		file.Close()
		return
	}

	prices := make([]float64, len(lines))

	for lineIndex, line := range lines {
		floatPrice, err := strconv.ParseFloat(line, 64)

		if err != nil {
			fmt.Printf("Erreur de conversion du prix à la ligne %d : %s\n", lineIndex+1, err)
			file.Close()
			return
		}

		prices[lineIndex] = floatPrice
	}

	job.InputPrices = prices
}

// Process calcule les prix TTC pour chaque prix HT de InputPrices
// en appliquant la formule : prixTTC = prix * (1 + TaxRate).
// Les résultats sont stockés dans TaxIncludedPrices avec le prix HT (formaté à 2 décimales) comme clé.
func (job TaxIncludedPriceJob) Process() {
	job.LoadData()

	result := make(map[string]float64)
	for _, price := range job.InputPrices {
		// clé : prix HT formaté à 2 décimales pour garantir l'unicité et la lisibilité
		result[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
	}

	fmt.Println(result)
}

// NewTaxIncludedPriceJob crée et retourne un nouveau TaxIncludedPriceJob
// avec un jeu de prix HT par défaut et le taux de taxe fourni en paramètre.
func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30}, // prix HT par défaut
		TaxRate:     taxRate,
	}
}
