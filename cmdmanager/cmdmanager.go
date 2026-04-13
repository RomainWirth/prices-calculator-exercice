package cmdmanager

import "fmt"

type CMDManager struct{}

func (cmd CMDManager) ReadLines() ([]string, error) {
	fmt.Println("S'il vous plaît, entrez les prix HT (un par ligne), puis confirmez chaque prix avec ENTER")

	var prices []string

	for {
		var price string
		fmt.Print("Prix HT : ")
		fmt.Scan(&price)

		if price == "0" {
			break
		}

		prices = append(prices, price)
	}

	return prices, nil
}

func (cmd CMDManager) WriteResult(data any) error {
	fmt.Println(data)
	return nil
}

func New() CMDManager {
	return CMDManager{}
}
