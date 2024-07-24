package main

import "fmt"

// Executado na hora da compilação
// antes de chamar a função main
// executa uma vez
func init() {
	fmt.Println("Inicializando...")
}

func main() {
	fmt.Println("Main...")
}