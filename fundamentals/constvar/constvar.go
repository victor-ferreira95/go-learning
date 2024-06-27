package main

import (
	"fmt"
	// se o nome for grande pode referenciar com algo na frente
	// se colocar o _ o pacote nao sera removido
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 //(tipo float65) foi inferido pelo copilador

	// forma reduzida de criar uma var
	// cria e atribui
	area := PI * m.Pow(raio, 2)

	fmt.Println("A area da circunferência é", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 2
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false

	fmt.Println(e, f)

	g, h, i := 2, false, "epa!"

	fmt.Println(g, h, i)
}
