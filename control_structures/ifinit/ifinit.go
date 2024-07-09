package main

import (
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return r.Intn(10)
}

func main() {
	// inicializa uma variável acessível apenas dentro do bloco if
	if i := numeroAleatorio(); i > 5 { // tb suporta no switch
		println("Ganhou!!!")
	} else {
		println("Perdeu :(")
	}

}
