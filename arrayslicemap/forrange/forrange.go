package main

import "fmt"

func main() {
	// array com tamanho definido na inicialização
	numeros := [...]int{1, 2, 3, 4, 5} // copilador conta!


	// equivalente ao foreach
	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i, numero)
	}

	for _, num := range numeros {
		fmt.Println(num)
	}
}
