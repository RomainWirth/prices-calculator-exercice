// Package prices fournit les outils pour calculer des prix TTC
// à partir d'un taux de taxe et d'une liste de prix HT.
package prices

import (
	"bufio" // lecture ligne par ligne du fichier
	"fmt"
	"os"      // accès au système de fichiers
	"strconv" // conversion des lignes texte en float64
)

// TaxIncludedPriceJob représente un travail de calcul de prix TTC.
// Il regroupe le taux de taxe, les prix HT en entrée et la map des prix TTC calculés.
type TaxIncludedPriceJob struct {
	TaxRate           float64            // taux de taxe à appliquer (ex: 0.2 pour 20%)
	InputPrices       []float64          // liste des prix HT à traiter
	TaxIncludedPrices map[string]float64 // résultats : clé = prix HT formaté, valeur = prix TTC
}

// LoadData lit le fichier "prices.txt" ligne par ligne et charge les prix HT dans job.InputPrices.
// Chaque ligne du fichier doit contenir un prix numérique (ex: "9.99").
// Le récepteur est un pointeur (*TaxIncludedPriceJob) pour que la modification de job.InputPrices
// soit visible en dehors de la méthode.
// En cas d'erreur d'ouverture, de lecture ou de conversion, un message est affiché et la fonction retourne.
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

	// allocation d'une slice de float64 de la même taille que le nombre de lignes lues
	prices := make([]float64, len(lines))

	for lineIndex, line := range lines {
		// conversion de la ligne texte en float64 (précision 64 bits)
		floatPrice, err := strconv.ParseFloat(line, 64)

		if err != nil {
			// la ligne ne contient pas un nombre valide : on ferme le fichier et on abandonne
			fmt.Printf("Erreur de conversion du prix à la ligne %d : %s\n", lineIndex+1, err)
			file.Close()
			return
		}

		prices[lineIndex] = floatPrice
	}

	// mise à jour du job avec les prix lus depuis le fichier (possible grâce au récepteur pointeur)
	job.InputPrices = prices
}

// Process charge les prix HT depuis le fichier, puis calcule les prix TTC pour chacun
// en appliquant la formule : prixTTC = prix * (1 + TaxRate).
// Le récepteur est un pointeur pour permettre à LoadData de mettre à jour job.InputPrices.
// Les résultats sont affichés sous forme de map (clé = prix HT à 2 décimales, valeur = prix TTC).
func (job *TaxIncludedPriceJob) Process() {
	// chargement des prix HT depuis prices.txt avant le calcul
	job.LoadData()

	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		// clé : prix HT formaté à 2 décimales pour garantir l'unicité et la lisibilité
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
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
