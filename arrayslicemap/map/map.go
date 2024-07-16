package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[12345678978] = "maria"
	aprovados[12345678990] = "pedro"
	aprovados[12345678908] = "Ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[12345678908])

	// deletar um elemento do map
	delete(aprovados, 12345678908)
	
	fmt.Println(aprovados[12345678908])
}
