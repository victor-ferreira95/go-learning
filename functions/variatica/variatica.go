package main

import "fmt"

// ... pode receber parâmetros variáveis
func media(numero ...float64) float64 {
	// ... é um slice
	total := 0.0
	for _, num := range numero {
		total += num
	}
	return total / float64(len(numero))
}

func main() {
	fmt.Printf("Média: %.2f", media(7.7, 8.1, 5.9, 9.9))
}
