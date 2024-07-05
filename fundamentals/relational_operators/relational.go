package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("igual:", "banana" == "banana")
	fmt.Println("diferente:", 3 != 2)
	fmt.Println("menor:", 3 < 2)
	fmt.Println("maior:", 3 > 2)
	fmt.Println("menor ou igual:", 3 <= 2)
	fmt.Println("maior ou igual:", 3 >= 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)
	fmt.Println("Datas:", d1 == d2)
	fmt.Println("Datas:", d1.Equal(d2))

	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"João"}
	p2 := Pessoa{"João"}
	fmt.Println("Pessoas:", p1 == p2)

}
