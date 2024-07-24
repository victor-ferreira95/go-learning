package main

// funções anonimas

var soma = func(a, b int) int {
	return a + b
}

func main() {
	println(soma(1, 2))

	// local
	sub := func(a, b int) int {
		return a - b
	}

	println(sub(5, 2))
}