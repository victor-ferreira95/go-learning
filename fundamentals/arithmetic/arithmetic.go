package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("soma", a+b)
	fmt.Println("subtração", a-b)
	fmt.Println("divisão", a/b)
	fmt.Println("multiplicação", a*b)
	fmt.Println("resto", a%b)

	// bitwise
	// operações binarias
	fmt.Println("E", a&b)   // 11 & 10 = 10
	fmt.Println("Ou", a|b)  // 11 | 10 = 11
	fmt.Println("Xor", a^b) // 11 ^ 10 = 01

	c := 3.0
	d := 2.0

	// outras operações usando math...
	fmt.Println("maior", math.Max(float64(a), float64(b)))
	fmt.Println("menor", math.Min(c, d))
	fmt.Println("exponenciação", math.Pow(c, d))
}
