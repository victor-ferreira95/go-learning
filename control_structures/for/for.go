package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1

	for i <= 10 { // não tem while em go, mas tem for
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	fmt.Println("\nMisturando...")

	for i = 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Par")
		} else {
			fmt.Println("Impar")
		}
	}

	for {
		// laco infinito
		fmt.Println("Laço infinito...")
		time.Sleep(time.Second) // pausa de 1 segundo
	}

	// veremos o foreach
}
