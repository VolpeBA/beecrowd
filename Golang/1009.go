package main

import (
	"fmt"
)

func main() {

	var nome string
	var salario_fixo, total_vendas float64

	fmt.Scan(&nome)
	fmt.Scan(&salario_fixo)
	fmt.Scan(&total_vendas)

	salario := salario_fixo + (total_vendas * 0.15)

	fmt.Printf("TOTAL = R$ %.2f\n", salario)
}
