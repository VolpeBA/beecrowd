package main

import (
	"fmt"
)

func main() {

	var valor_peca1, valor_peca2, cod_peca1, num_peca1, cod_peca2, num_peca2 float64

	fmt.Scan(&cod_peca1, &num_peca1, &valor_peca1)
	fmt.Scan(&cod_peca2, &num_peca2, &valor_peca2)

	valor_total := (valor_peca1 * num_peca1) + (valor_peca2 * num_peca2)

	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", valor_total)
}
