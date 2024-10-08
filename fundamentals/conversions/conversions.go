package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// cuidado com o casting, código tabela ASC
	fmt.Println("Teste " + string(123))

	// int para string
	fmt.Println("Teste " + strconv.Itoa(1234))

	num, _ := strconv.Atoi("123")
	// string para int
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro")
	}
}
