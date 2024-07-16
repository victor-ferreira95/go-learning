package main

import (
	"fmt"
)

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 12345.6,
			"Guga Pereira":   12345.7,
		},
		"J": {
			"José João": 11223.45,
		},
		"P": {
			"Pedro Junior": 1200.0,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
