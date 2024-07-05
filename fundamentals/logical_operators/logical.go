package main

import "fmt"

func comprar(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2
	comprarTv32 := trab1 != trab2 // ou exclusivo
	comprouSorvete := trab1 || trab2

	return comprarTv50, comprarTv32, comprouSorvete
}

func main() {
	tv50, tv32, sorvete := comprar(true, true)
	fmt.Printf("tv50: %t, tv32: %t, sorvete: %t, saudável: %t", tv50, tv32, sorvete, !sorvete)
}
