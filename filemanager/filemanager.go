package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

func ReadLines(path string) ([]string, error) {
	// ouverture du fichier contenant les prix HT
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New("Erreur d'ouverture du fichier")
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
		file.Close()
		return nil, errors.New("Erreur lors de la lecture du fichier")
	}

	file.Close()
	return lines, nil
}

func WriteJSON(path string, data interface{}) error {
	file, err := os.Create(path)
	if err != nil {
		return errors.New("Erreur de création du fichier")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		file.Close()
		return errors.New("Erreur d'écriture dans le fichier")
	}

	file.Close()
	return nil
}
