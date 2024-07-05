package main

import "fmt"

// referencia de memoria

func main() {
	i := 1

	var p *int = nil //nulo
	p = &i // pegando o endereço da variável
	*p++   // desreferenciando (pegando o valor)
	i++

	// Go não tem aritmética de ponteiros
	// p++

	fmt.Println(p, *p, i, &i)
}
