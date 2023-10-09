package main

import (
	"fmt"
)

func main() {

	var km, valor_combustivel float64

	fmt.Scan(&km)
	fmt.Scan(&valor_combustivel)

	consumo := km / valor_combustivel

	fmt.Printf("%.3f km/l\n", consumo)
}
