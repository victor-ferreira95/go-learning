package main

func notaParaConceito(n float64) string {

	switch {
	case n >= 9 && n <= 10:
		return "A"
	case n >= 8 && n < 9:
		return "B"
	case n >= 5 && n < 8:
		return "C"
	case n >= 3 && n < 5:
		return "D"
	default:
		return "E"

	}
}

func main() {
	println(notaParaConceito(9.8))
	println(notaParaConceito(8.5))
	println(notaParaConceito(6.9))
	println(notaParaConceito(4.1))
	println(notaParaConceito(2.3))
}
