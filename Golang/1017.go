package main

import (
	"fmt"
)

func main() {

	var tempo_gasto, velocidade float64

	fmt.Scan(&tempo_gasto)
	fmt.Scan(&velocidade)

	quantidade_litros := ((velocidade / 12.0) * tempo_gasto)

	fmt.Printf("%.3f\n", quantidade_litros)
}
