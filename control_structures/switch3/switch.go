package main

// tipo genérico
func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "função"
	default:
		return "nenhum"
	}
}

func main() {
	// Funções são tipos
	println(tipo(func() {}))
	println(tipo(10))
	println(tipo(10.0))
	println(tipo("texto"))
	println(tipo(true))
	println(tipo(nil))
}
